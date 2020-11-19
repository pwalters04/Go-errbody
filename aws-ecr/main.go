package main

import (
	ecr_operations "aws-ecr/ecr-operations"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ecr"
)
///srcImage="$ECR_SRC_REG_ID.dkr.ecr.$ECR_SRC_REGION.${AWS::URLSuffix}/$ECR_REPO_NAME:$ECR_REPO_TAG"
func main (){


	var Found bool
	ecrService :=ecr_operations.ECRService("Admin-dev","us-east-2")
	repoNamesResults,_:=ecr_operations.ListECRRepos("Admin-dev","us-east-2")
	loginToken := ecr_operations.ECR_Login("Admin-dev","us-east-2")
	regID := "513877516433"
	region := "us-east-2"
	repoName := "sample"
    srcImage :=regID +".dkr.ecr.region.amazonaws.com/"+repoName+":latest"
    targetImage:= regID +".dkr.ecr.us-east-1.amazonaws.com/"+repoName+":latest"

 for _,theList := range repoNamesResults{
 	 if theList == "TestRepo"{
		  Found = true
	  }else{
	  		Found = false
	 	}
	}
// Create Repo
	if Found == false {
		fmt.Println("Creating Repo....")

		createInput := &ecr.CreateRepositoryInput{
			RepositoryName: aws.String("testing-repo/testing"),
		}

		createResult, err := ecrService.CreateRepository(createInput)

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case ecr.ErrCodeServerException:
					fmt.Println(ecr.ErrCodeServerException, aerr.Error())
				case ecr.ErrCodeInvalidParameterException:
					fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
				case ecr.ErrCodeInvalidTagParameterException:
					fmt.Println(ecr.ErrCodeInvalidTagParameterException, aerr.Error())
				case ecr.ErrCodeTooManyTagsException:
					fmt.Println(ecr.ErrCodeTooManyTagsException, aerr.Error())
				case ecr.ErrCodeRepositoryAlreadyExistsException:
					fmt.Println(ecr.ErrCodeRepositoryAlreadyExistsException, aerr.Error())
				case ecr.ErrCodeLimitExceededException:
					fmt.Println(ecr.ErrCodeLimitExceededException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {

				fmt.Println(err.Error())
			}
			return
		}

		fmt.Println(createResult)
	}
}





