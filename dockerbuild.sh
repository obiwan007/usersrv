docker build -t obiwan007/gqlsrv -f ./gql_Dockerfile . 
docker push obiwan007/gqlsrv

docker build -t obiwan007/usersrv -f ./user_Dockerfile . 
docker push obiwan007/usersrv