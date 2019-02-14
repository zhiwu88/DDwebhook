package transformer

import (
    "bytes"
    "fmt"
    model "model/ddmodel"
)

func TransformToMarkdown(notification model.Notification) (markdown *model.DingTalkMarkdown, robotURL string, err error) {

    status := notification.Status
    commonAnnotations := notification.CommonAnnotations
    robotURL = commonAnnotations["dingtalkRobot"]

    var buffer bytes.Buffer
    buffer.WriteString(fmt.Sprintf("### [%s] \n", status))
    buffer.WriteString(fmt.Sprintf("#### Summary: %s \n", commonAnnotations["summary"]))

    for _, alert := range notification.Alerts {
        labels := alert.Labels
        alert_annotations := alert.Annotations
        buffer.WriteString(fmt.Sprintf("\n> %s \n", alert_annotations["description"]))
        buffer.WriteString(fmt.Sprintf("\n> 开始时间:%s \n", alert.StartsAt.Format("2006-01-02 15:04:05")))

        buffer.WriteString(fmt.Sprintf("\n #### Labels \n"))
        for labelName, labelValue := range labels {
            buffer.WriteString(fmt.Sprintf("\n> %s=\"%s\" \n", labelName, labelValue))
        }

    }

    markdown = &model.DingTalkMarkdown {
        MsgType: "markdown",
        Markdown: &model.Markdown {
            Title: fmt.Sprintf("[%s] %s", status, commonAnnotations["summary"]),
            Text: buffer.String(),
        },
        At: &model.At {
            IsAtAll: false,
        },
    }

    return

}

