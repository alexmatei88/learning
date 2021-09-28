package main;

import "fmt"
import "os"
import "strings"
import "io/ioutil"
import "net/http"
import "encoding/json"



func main(){
	fmt.Println("---------------- First we read the name of the backup server from arguments -----------------")

	backupsServerName := os.Args[1];
	token := os.Args[2];
	mwpToken := os.Args[3];
	fmt.Println("Server is: " , string(backupsServerName));
	fmt.Println("Token is: " , string(token));
	fmt.Println("MWP Token is: ", string(mwpToken));

	endpointUrlTemplate := "https://<servername>.sucuri.net/mwp-export.php?token=<TOKEN>&action=users"
	mwpEndpointURLTemplate := "/api/v1/user/{$user_id}/backup?token=<TOKEN>";

	endpointURL := strings.Replace(endpointUrlTemplate,"<servername>", os.Args[1],1);
	endpointURL = strings.Replace(endpointURL,"<TOKEN>",token,1);

	mwpEndpointURL := strings.Replace(mwpEndpointURLTemplate,"<TOKEN>",mwpToken,1);

	fmt.Println("Endpoint URL este ", string(endpointURL));
	fmt.Println("MWP Endpoint URL este ", string(mwpEndpointURL));

	retrieveAccLit(endpointURL);
}

func retrieveAccLit(url string){
	fmt.Println("---------------- Going to call the following endpoint" , string(url));

	response , err := http.Get(url);
	if(err!=nil){
		fmt.Println(err.Error());
	}

	responseData, err := ioutil.ReadAll(response.Body);

	if(err!=nil){
		fmt.Println(err.Error());
	}

	fmt.Println(string(responseData));
	unmarshallResponse(responseData);
}

func unmarshallResponse(responseData []byte){
	responseObjects := make([]AccountDetails,0);
	json.Unmarshal(responseData,&responseObjects);

	for i:= range responseObjects{
	fmt.Println("Account: " , string(responseObjects[i].Email));
	checkStatusWithMWP(responseObjects[i]);
	}
	
}

func checkStatusWithMWP(account AccountDetails){
	
}

type Response struct {
	Collection []AccountDetails 
}

type AccountDetails struct{
	Email string
	Alt_email string 
	Locale string 
	Plan string 
	Account_ssh_key string 
}
