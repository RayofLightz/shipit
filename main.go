package main

import (
        "fmt"
        "net"
        "os"
        "flag"
        "bufio"
)

func sendRecord(jsonLine string, addr string) error{
        fmt.Println("sending")
        conn, err := net.Dial("tcp", addr)
        if err != nil{
                return err
        }
        bytes_wrote, err := conn.Write([]byte(jsonLine))
        if err != nil{
                return err
        }
        fmt.Println("Bytes wrote ",bytes_wrote)
        return nil
}

func sendFile(filename string, addr string) error{
        file, err := os.Open(filename)
        fmt.Println("reading file")
        if err != nil{
                return err
        }
        scanner := bufio.NewScanner(file)
        scanner.Split(bufio.ScanLines)
        for scanner.Scan(){
                err = sendRecord(scanner.Text(), addr)
                if err != nil{
                        return err
                }
        }
        return nil
}

func main(){
        //Program flags
        var file string
        var addr string
        flag.StringVar(&file,"FileName","","log file")
        flag.StringVar(&addr,"Address","","address of ingestion ex 127.0.0.1:8080")
        flag.Parse()
        if file == "" || addr == ""{
                fmt.Println("Usage $./shipit -FileName <logfile> -Address <ingestion addr>")
                return
        }
        err := sendFile(file,addr)
        if err != nil{
                fmt.Println(err)
        }
}
