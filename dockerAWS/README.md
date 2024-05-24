# docker & AWS

Demo for deploying Go app on AWS EC2 with docker

Create AWS account (its free)
create EC2 instance : its PC on cloud

`Steps` :
EC2 >> EC2 virtual servers in cloud >> Instances >> Launch instance >> name & tag of instance >> Application & OS image >> select from available (look for free ones) >> instance type (CPU) t2.micro >> create key pair (for login & secure connection between your PC & instance) & download it >> Config Storage (hard disk) >> launch instance
(skips setting steps which are not mentioned here)

Select instance from dashboard >> click connect >> follow steps

Go to working dir on PC >> copy & run cmds on terminal (run all cmds with sudo)
wallah now you are connected with Cloud PC

Now to deploy on AWS :
git clone repo
sudo install docker.io
sudo snap install docker
sudo chmod 777 /var/run/docker.sock  
(777 -permission to read write etc)

Now build docker & continue with docker cmds (refer `docker1`)

Now on browser :
localhost:8000 won't work you need Instance IP for this & allow ports
Go to your instace >> security >> click on Security group link >> inbound rule >> edit >> add rule >> all traffic >> anywhere-IPv4 >> save rules
Go to your instace >> details >> copy Public IPv4 address >> paste on browser >> add port
ex : 18.222.21.242:8000
      IP address  :port





