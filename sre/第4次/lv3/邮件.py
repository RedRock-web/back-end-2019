import smtplib
from email.mime.text import MIMEText

sender = "sdfisd@163.com"   #����������
passwd = "123456"          #�������ʼ���Ȩ��
receivers = "12345@qq.com"      #�ռ�������

subject = "python�ʼ�����"  #����
content = "hello"        #��ͨ����
# content = """           #html��ʽ����
# <p>Hello World!</p>
# """

msg = MIMEText(content, "plain", "utf-8")

msg["Subject"] = subject
msg["From"] = sender
msg["To"] = receivers



try:
    s = smtplib.SMTP_SSL("smtp.163.com", 465)    #����ssl������������smtp
    s.login(sender, passwd)         #��¼
    s.sendmail(sender, receivers, msg.as_string())        #����
    print("ok!")

except Exception:
    print("no!")






