package utils

import (
  "fmt"
  "net"
  "net/smtp"
  "crypto/tls"

  // "github.com/scorredoira/email"
  "github.com/matcornic/hermes"

  "github.com/hiyali/katip-be/config"
)

func getConfiguredEmail() hermes.Hermes {
  cf := config.Get()

  return hermes.Hermes{
    Product: hermes.Product{
      Name: cf.App.Name,
      Link: cf.App.Link,
      // Logo: "https://katip.hiyali.org/assets/logo.png",
      Copyright: cf.AppCopyright(),
    },
  }
}

func SendRegisterConfirmEmail(userEmail string, userName string, token string) (err error) {
  cf := config.Get()

  // body
  hermesEmail := hermes.Email{
    Body: hermes.Body{
      Name: userName,
      Intros: []string{
        "You have received this email because a register request for " + cf.App.Name + " account was received.",
      },
      Actions: []hermes.Action{
        {
          Instructions: "Click the button below to finish register request:",
          Button: hermes.Button{
            Color: "#22BC66",
            Text:  "Confirm Register",
            Link:  cf.App.Link + "api/register-confirm?token=" + token,
          },
        },
      },
      Outros: []string{
        "If you did not request register " + cf.App.Name + ", no further action is required on your part.",
      },
      Signature: "Thanks",
    },
  }

  h := getConfiguredEmail()
  emailBody, err := h.GenerateHTML(hermesEmail)
  if err != nil {
    return err
  }

  /*
  m := email.NewHTMLMessage("Congratulations! You registered the Katip.", emailBody)
  if err := m.Attach("email.go"); err != nil {
    log.Fatal(err)
  }
  */

  // header
  header := make(map[string]string)
  header["From"] = cf.Email.Name + "<" + cf.Email.Address + ">"
  header["To"] = userEmail
  header["Subject"] = "You are bing register " + cf.App.Name
  header["Content-Type"] = "text/html; charset=UTF-8"

  // message
  message := ""
  for k, v := range header {
    message += fmt.Sprintf("%s: %s\r\n", k, v)
  }
  message += "\r\n" + emailBody

  // send it
  auth := smtp.PlainAuth("", cf.Email.Address, cf.Email.Password, cf.Email.ServerHost)
  if err = SendMailUsingTLS(cf.EmailHostName(), auth, cf.Email.Address, []string{userEmail}, []byte(message)); err != nil {
    return
  }
  return nil
}

// return a smtp client
func Dial(addr string) (*smtp.Client, error) {
  conn, err := tls.Dial("tcp", addr, nil)
  if err != nil {
    return nil, err
  }
  host, _, _ := net.SplitHostPort(addr)
  return smtp.NewClient(conn, host)
}

func SendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {
  c, err := Dial(addr)
  if err != nil {
    return err
  }
  defer c.Close()

  if auth != nil {
    if ok, _ := c.Extension("AUTH"); ok {
      if err = c.Auth(auth); err != nil {
        return err
      }
    }
  }

  if err = c.Mail(from); err != nil {
    return err
  }

  for _, addr := range to {
    if err = c.Rcpt(addr); err != nil {
      return err
    }
  }

  w, err := c.Data()
  if err != nil {
    return err
  }

  _, err = w.Write(msg)
  if err != nil {
    return err
  }

  err = w.Close()
  if err != nil {
    return err
  }

  return c.Quit()
}
