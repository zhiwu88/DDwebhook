package notifier

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    model "model/ddmodel"
    "transformer/ddtransformer"
)

func Send(notification model.Notification, defaultRobot string) (err error) {
    
    markdown, robotURL, err := transformer.TransformToMarkdown(notification)
    
    fmt.Printf("%+v\n", markdown)

    if err != nil {
        return
    }

    data, err := json.Marshal(markdown)
    if err != nil {
        return
    }

    fmt.Println(string(data))

    var dingTalkRobotURL string
    if robotURL != "" {
        dingTalkRobotURL = robotURL
    } else {
        dingTalkRobotURL = defaultRobot
    }

    req, err := http.NewRequest(
        "POST",
        dingTalkRobotURL,
        bytes.NewBuffer(data))

    if err != nil {
        return
    }

    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        return
    }

    defer resp.Body.Close()
    fmt.Println("response Status:", resp.Status)
    fmt.Println("respose Headers:", resp.Header)

    return

}
