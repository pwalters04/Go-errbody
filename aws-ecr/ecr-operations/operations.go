package ecr_operations

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"log"
)

func ECRService(profileName string, region string ) *ecr.ECR{
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile:profileName,
		Config: aws.Config{
			Region: aws.String(region),
		}})
	if err != nil {
		fmt.Println("[Error Session]")
		log.Fatal()
	}

	ecrService := ecr.New(sess)
	return ecrService
}

func ListECRRepos(profileName string, region string )([]string, error){
	var repoNames [] string
	ecrSrv := ECRService(profileName,region)
	input := &ecr.DescribeRepositoriesInput{}
	results,err := ecrSrv.DescribeRepositories(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil,nil
	}

	for _,ecrRepo := range results.Repositories{
		repoNames = append(repoNames,aws.StringValue(ecrRepo.RepositoryName))
	}
	return repoNames,nil
}

func ECR_Login(profileName string, region string) *ecr.GetAuthorizationTokenOutput {
	ecrSrv := ECRService(profileName,region)
	input := &ecr.GetAuthorizationTokenInput{}

	result, err := ecrSrv.GetAuthorizationToken(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil
	}

	return result
}

//func GetImage(profileName string, region string, imageName string)  {
//	ecrSrv := ECRService(profileName,region)
//	input := &ecr.ListImagesInput{
//		RepositoryName: aws.String(imageName),
//	}
//
//	result, err := ecrSrv.ListImages(input)
//	if err != nil {
//		if aerr, ok := err.(awserr.Error); ok {
//			switch aerr.Code() {
//			case ecr.ErrCodeServerException:
//				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
//			case ecr.ErrCodeInvalidParameterException:
//				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
//			case ecr.ErrCodeRepositoryNotFoundException:
//				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
//			default:
//				fmt.Println(aerr.Error())
//			}
//		} else {
//			// Print the error, cast err to awserr.Error to get the Code and
//			// Message from an error.
//			fmt.Println(err.Error())
//		}
//		return
//	}
//
//	fmt.Println(result)
//	result.ImageIds
//
//}