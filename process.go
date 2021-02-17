package password

import ("math/rand"
"fmt"
 "strings"
"strconv"
"flag"
"time"
"os")

func RunFirst(){
	opt,pwd := getOperator()
    rand.Seed(time.Now().UnixNano())
    var chk []string
	var num int
    if opt == "check"{ 
        num, chk = chkpwd(pwd)
        fmt.Println(pwd, chk )
    	fmt.Println("Corrections needed in password",num)
    }else{ 
        if opt == "generate"{
		   p := genpwd() 
		   fmt.Println("passowrd",p)
        }
	}
	
} 

//Get the operation to be performed check/generate 
func getOperator()(string,string){
    fs := flag.NewFlagSet("Flags", flag.ExitOnError)
    var args []string 
    args = os.Args[1:]
    fs.Parse(args) 
    opt := fs.Arg(0)
    pwd := fs.Arg(1)
	return opt,pwd
}

//check the given password for length, spl character, contains alpahbet etc
func chkpwd(pwd string)(int, []string){
    var chk string
    var chks []string
    var num int
    num = 1
     if len(pwd) < 8{ 
         chk = "Password length is less, First increase to min 8"
         chks= append(chks,chk)
        return num,chks 
     }else{ 
         if len(pwd) < 15{ 
             num, chks = evaluatepwd(pwd) 
             return num,chks
         }else if len(pwd) < 20{ 
             num, chks = evaluatepwd(pwd)
             return num,chks
         }else{ 
            chk = "Password length is more than required, reduce to max 20"
            chks= append(chks,chk)
             return num,chks
         }
     } 
}

func evaluatepwd(pwd string)(int,[]string){ 
    var chk string
    var chks []string
	change := 0
    if !strings.ContainsAny(pwd,"abcdefghijklmnopqrstuvwxyz"){ 
        change += 1 
        chk = "No Lower case letter"
        chks = append(chks,chk)
    } 
    if !strings.ContainsAny(pwd,"ABCDEFGHIJKLMNOPQRSTUVWXYZ"){ 
        chk = "No Upper case letter" 
        change += 1
        chks = append(chks,chk)
    } 
    if !strings.ContainsAny(pwd,"0123456789"){ 
        chk = "No Number" 
        change += 1
        chks = append(chks,chk)
    } 
    if !strings.ContainsAny(pwd,"$#@!&*^%"){ 
        chk = "No Special Character" 
        change += 1
        chks = append(chks,chk)
    }
    same := false
    prev:=pwd[0] 
    for i:=1;i<len(pwd);i++{
         if pwd[i] == prev{ 
             if same{ 
                 same = false
                 chk = "3 Con characters are same"
                 chks = append(chks,chk)
                 change += 1
             }else{
                 same = true
             }
         }else{
             same = false
         }
         prev = pwd[i]
    }

    return change,chks 
} 

//generates a new pwd with length 15

func genpwd() string{
	 a :=[]string{}
	 //a = {} 
	 var spl,upper,lower string
	 spl = "$#@!&*^%"
	 upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	 lower = "abcdefghijklmnopqrstuvwxyz"
     up := randomnumber(15)  //UC position 
     np := randomnumber(15) //number position 
     sp := randomnumber(15) //specialcharacter postion 
     
     for i := 0;i<15;i++{ 
         switch i{
             case np:
                a = append(a,strconv.Itoa(randomnumber(10)))
             case sp: 
                a = append(a,string(spl[randomnumber(8)]))
             case up : 
                a = append(a,string(upper[randomnumber(8)]))
             default: 
             a =  append(a,string(lower[randomnumber(26)]))
        }
  
	 }  
	 s := strings.Join(a,"")
	 return s           
} 



func randomnumber(n int)int{
    n1 := rand.Intn(n)
    return n1 
}