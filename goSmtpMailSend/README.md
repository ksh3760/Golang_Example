This is a STMP email send package example <br/>
The Example used google smtp server <br/>
Write config.conf & mailSet.json files <br/>


config.conf <br/>
```
{
    "g_mail_id" : "yourGmailHere@gmail.com",
    "g_mail_pw" : "yourPasswordHere",
    "g_smtp_host" : "smtp.gmail.com",
    "g_smtp_port" : "587"
}
```

mailSet.json <br/>
```
{
    "title" : "example Mail title",
    "content" : "example Mail content",
    "recipient" : [
        "recipientEmailHere@gmail.com",
        "john.doe@gmail.com"
    ]
}
```
