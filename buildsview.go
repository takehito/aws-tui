package main

import (
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/rivo/tview"
)

func createBuildsView(ids []string) tview.Primitive {
	list := tview.NewList().ShowSecondaryText(false)
	list.SetBorder(true)

	for _, v := range ids {
		list.AddItem(v, "", []rune(v)[0], func() {})
	}

	builddescs, err := getBuildsDescription(ids)
	if err != nil {
		panic(err)
	}

	description := tview.NewTextView()
	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		build := builddescs[index]
		description.SetText(fmtBuild(build))
	})
	description.SetText(fmtBuild(builddescs[0]))

	return tview.NewFlex().
		AddItem(list, 0, 1, true).
		AddItem(description, 0, 2, false)
}

func fmtBuild(b types.Build) string {
	sb := strings.Builder{}

	sb.WriteString("Arn:\n")
	if b.Arn != nil {
		sb.WriteString(*b.Arn)
	}
	sb.WriteString("\n")

	sb.WriteString("Artifacts:\n")
	if b.Artifacts != nil {
		sb.WriteString("BuildArtifacts")
	}
	sb.WriteString("\n")

	sb.WriteString("BuildBatchArn:\n")
	if b.BuildBatchArn != nil {
		sb.WriteString(*b.BuildBatchArn)
	}
	sb.WriteString("\n")

	sb.WriteString("BuildComplete:\n")
	if b.BuildComplete {
		sb.WriteString("complete")
	} else {
		sb.WriteString("in proggress")
	}
	sb.WriteString("\n")

	sb.WriteString("BuildNumber:\n")
	sb.WriteString(strconv.Itoa(int(*b.BuildNumber)))
	sb.WriteString("\n")

	sb.WriteString("BuildStatus StatusType:\n")
	sb.WriteString(string(b.BuildStatus))
	sb.WriteString("\n")

	sb.WriteString("Cache: ")
	if b.Cache != nil {
		sb.WriteString("*ProjectCache")
	}
	sb.WriteString("\n")

	sb.WriteString("CurrentPhase:\n")
	if b.CurrentPhase != nil {
		sb.WriteString(*b.CurrentPhase)
	}
	sb.WriteString("\n")

	sb.WriteString("DebugSession:\n")
	if b.DebugSession != nil {
		sb.WriteString("*DebugSession")
	}
	sb.WriteString("\n")

	sb.WriteString("EncryptionKey:\n")
	if b.EncryptionKey != nil {
		sb.WriteString(*b.EncryptionKey)
	}
	sb.WriteString("\n")

	sb.WriteString("EndTime:\n")
	if b.EndTime != nil {
		sb.WriteString(b.EndTime.String())
	}
	sb.WriteString("\n")

	sb.WriteString("Environment:\n")
	if b.Environment != nil {
		sb.WriteString("*ProjectEnvironment")
	}
	sb.WriteString("\n")

	sb.WriteString("ExportedEnvironmentVariables:\n")
	if b.ExportedEnvironmentVariables != nil {
		sb.WriteString("[]ExportedEnvironmentVariable")
	}
	sb.WriteString("\n")

	sb.WriteString("FileSystemLocations:\n")
	if b.FileSystemLocations != nil {
		sb.WriteString("[]ProjectFileSystemLocation")
	}
	sb.WriteString("\n")

	sb.WriteString("Id:\n")
	if b.Id != nil {
		sb.WriteString(*b.Id)
	}
	sb.WriteString("\n")

	sb.WriteString("Initiator:\n")
	if b.Initiator != nil {
		sb.WriteString(*b.Initiator)
	}
	sb.WriteString("\n")

	sb.WriteString("Logs:\n")
	if b.Logs != nil {
		sb.WriteString("*LogsLocation")
	}
	sb.WriteString("\n")

	sb.WriteString("NetworkInterface:\n")
	if b.NetworkInterface != nil {
		sb.WriteString("*NetworkInterface")
	}
	sb.WriteString("\n")

	sb.WriteString("Phases:\n")
	if b.Phases != nil {
		sb.WriteString("[]BuildPhase")
	}
	sb.WriteString("\n")

	sb.WriteString("ProjectName:\n")
	if b.ProjectName != nil {
		sb.WriteString(*b.ProjectName)
	}
	sb.WriteString("\n")

	sb.WriteString("QueuedTimeoutInMinutes:\n")
	if b.QueuedTimeoutInMinutes != nil {
		sb.WriteString("*int32")
	}
	sb.WriteString("\n")

	sb.WriteString("ReportArns:\n")
	sb.WriteString(strings.Join(b.ReportArns, "\n"))
	sb.WriteString("\n")

	sb.WriteString("ResolvedSourceVersion:\n")
	if b.ResolvedSourceVersion != nil {
		sb.WriteString(*b.ResolvedSourceVersion)
	}
	sb.WriteString("\n")

	sb.WriteString("SecondaryArtifacts:\n")
	for i, _ := range b.SecondaryArtifacts {
		sb.WriteString(strconv.Itoa(i))
		sb.WriteString("\n")
	}
	sb.WriteString("\n")

	sb.WriteString("SecondarySourceVersions:\n")
	for i, _ := range b.SecondarySourceVersions {
		sb.WriteString(strconv.Itoa(i))
		sb.WriteString("\n")
	}
	sb.WriteString("\n")

	sb.WriteString("SecondarySources:\n")
	for i, _ := range b.SecondarySources {
		sb.WriteString(strconv.Itoa(i))
		sb.WriteString("\n")
	}
	sb.WriteString("\n")

	sb.WriteString("ServiceRole:\n")
	if b.ServiceRole != nil {
		sb.WriteString(*b.ServiceRole)
	}
	sb.WriteString("\n")

	sb.WriteString("Source:\n")
	if b.Source != nil {
		sb.WriteString("*ProjectSource")
	}
	sb.WriteString("\n")

	sb.WriteString("SourceVersion:\n")
	if b.SourceVersion != nil {
		sb.WriteString(*b.SourceVersion)
	}
	sb.WriteString("\n")

	sb.WriteString("StartTime:\n")
	if b.StartTime != nil {
		sb.WriteString(b.StartTime.String())
	}
	sb.WriteString("\n")

	sb.WriteString("TimeoutInMinutes:\n")
	if b.TimeoutInMinutes != nil {
		sb.WriteString(strconv.Itoa(int(*b.TimeoutInMinutes)))
	}
	sb.WriteString("\n")

	sb.WriteString("VpcConfig:\n")
	if b.VpcConfig != nil {
		sb.WriteString("*VpcConfig")
	}
	sb.WriteString("\n")

	return sb.String()
}
