package db

import (
    "github.com/joho/godotenv"
    "log"
    supa "github.com/nedpals/supabase-go"
    "os"
)

var Supabase *supa.Client

func CreateClient(){
    err := godotenv.Load(".env")

    if err != nil{
        log.Fatal("Error loading the .env file")
    }

    url := os.Getenv("SUPABASE_URL")
    key := os.Getenv("SUPABASE_KEY")

    Supabase = supa.CreateClient(url, key)
}
