# SE_PreLabTest2_2

```bash
ls -al ~/.ssh
mkdir ~/.ssh 
ssh-keygen -t rsa 4096 -b -C “ramnarong5259@gmail.com” 
clip < ~/.ssh/id_rsa.pub 
```

```bash
git clone git@github.com:sut67/team05.git 
git clone https://github.com/Ramnarong06/SE_PreLabTest2_2.git
```

```bash
git init
```

cd backend
```bash
go mod init github.com/Ramnarong06/SE_PreLabTest2_2
go get -u gorm.io/gorm
go get -u github.com/asaskavich/govalidator
go get -u github.com/onsi/gomega
```

```bash
git status
git checkout -b main
git checkout -b issue-2
```
ปิด vs แล้วเปิดใหม่

```bash
git status
```

```bash
git add .
git commit -m "ข้อความ - close #2"
git push origin main
git push origin issue-2
```