import smtplib
from email.mime.text import MIMEText

sender = "sdfisd@163.com"   #发送人邮箱
passwd = "123456"          #发送人邮件授权码
receivers = "12345@qq.com"      #收件人邮箱

subject = "python邮件测试"  #主题
content = "hello"        #普通正文
# content = """           #html格式正文
# <p>Hello World!</p>
# """

msg = MIMEText(content, "plain", "utf-8")

msg["Subject"] = subject
msg["From"] = sender
msg["To"] = receivers



try:
    s = smtplib.SMTP_SSL("smtp.163.com", 465)    #创建ssl变量连接网易smtp
    s.login(sender, passwd)         #登录
    s.sendmail(sender, receivers, msg.as_string())        #发送
    print("ok!")

except Exception:
    print("no!")






