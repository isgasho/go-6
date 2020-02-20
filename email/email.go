/*
Package email implements the formatting of multipart RFC 2045 e-mail messages,
including headers, attachments, HTML email, and plain text.

Example:

    package main
    
    import (
        "net/mail"
        "os"
        
        "tawesoft.co.uk/go/email"
    )
    
    func main() {
        var eml = email.Message{
            From:  mail.Address{"Alan Turing", "turing.alan@example.org"},
            To:  []mail.Address{{"Grace Hopper", "amazing.grace@example.net"}},
            Bcc: []mail.Address{{"BCC1", "bcc1@example.net"}, {"BCC2", "bbc2@example.net"}},
            Subject: "Computer Science is Cool! ❤",
            Text: `This is a test email!`,
            Html: `<!DOCTYPE html><html lang="en"><body><p>This is a test email!</p></body></html>`,
            Attachments: []*email.Attachment{
                //email.FileAttachment("Entscheidungsproblem.pdf"),
                //email.FileAttachment("funny-cat-meme.png"),
            },
            Headers: mail.Header{
                "X-Category": []string{"newsletter", "marketing"},
            },
        }
        
        var err = eml.Print(os.Stdout)
        if err != nil { panic(err) }
    }

For license information, documentation, source code, support, links, etc. please see
https://tawesoft.co.uk/go/email

This module is part of https://tawesoft.co.uk/go
*/
package email // import "tawesoft.co.uk/go/email"

// SPDX-License-Identifier: MIT

// Code generated by (tawesoft.co.uk/go/) fluff.py: DO NOT EDIT.