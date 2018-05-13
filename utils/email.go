package utils

import (
  "log"
  "fmt"
  "time"
  "net"
  "net/smtp"
  "crypto/tls"

  // "github.com/scorredoira/email"
  "github.com/matcornic/hermes"
)

type (
  ProductConfig struct {
    Name  string
    Link  string
  }

  EmailConfig struct {
    Address    string
    Password   string
    ServerHost string
    ServerPort uint
  }
)

func (pc *ProductConfig) Copyright() string {
  return fmt.Sprintf("Copyright © %v %v. All rights reserved.", time.Now().Year(), pc.Name)
}

func (ec *EmailConfig) HostName() string {
  return fmt.Sprintf("%v:%v", ec.ServerHost, ec.ServerPort)
}

var prdConf ProductConfig
var emailConf EmailConfig

func init() {
  prdConf = ProductConfig{
    Name: "Katip",
    Link: "https://katip.hiyali.org/",
  }
  emailConf = EmailConfig{
    Address   : "katip-team@hiyali.org", // "salam.14@163.com",
    Password  : "non-secure",
    ServerHost: "smtp.hiyali.org", // "smtp.163.com",
    ServerPort: 994,
  }
}

func getConfiguredEmail() hermes.Hermes {
  return hermes.Hermes{
    Product: hermes.Product{
      Name: prdConf.Name,
      Link: prdConf.Link,
      // Logo: "http://katip.hiyali.org/assets/logo.png",
      Copyright: prdConf.Copyright(),
    },
  }
}

func SendRegisterEmail(userEmail string, userName string, userPassword string) (err error) {
  // body
  hermesEmail := hermes.Email{
    Body: hermes.Body{
      Name: userName,
      Intros: []string{
        "You have received this email because a password reset request for Katip account was received.",
      },
      Actions: []hermes.Action{
        {
          Instructions: "Click the button below to reset your password:",
          Button: hermes.Button{
            Color: "#DC4D2F",
            Text:  "Reset your password: " + userPassword,
            Link:  prdConf.Link + "reset-password?token=d9729feb74992cc3482b350163a1a010",
          },
        },
      },
      Outros: []string{
        "If you did not request a password reset, no further action is required on your part.",
      },
      Signature: "Thanks!",
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
  header["From"] = "Katip Team" + "<" + emailConf.Address + ">"
  header["To"] = userEmail
  header["Subject"] = "Congratulations! You registered the Katip product."
  header["Content-Type"] = "text/html; charset=UTF-8"

  // message
  message := ""
  for k, v := range header {
    message += fmt.Sprintf("%s: %s\r\n", k, v)
  }
  message += "\r\n" + emailBody

  // send it
  auth := smtp.PlainAuth("", emailConf.Address, emailConf.Password, emailConf.ServerHost)
  if err = SendMailUsingTLS(emailConf.HostName(), auth, emailConf.Address, []string{userEmail}, []byte(message)); err != nil {
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
    log.Println("SendMailUsingTLS Dial", err)
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
