migrate -database "postgresql://$POSTGRES_USER:$POSTGRES_PASS@$POSTGRES_HOST/$POSTGRES_DATABASE" -path migrations up