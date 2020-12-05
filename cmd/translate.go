package cmd

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	translate "cloud.google.com/go/translate/apiv3"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

var (
	srcLngCode      string
	dstLngCode      string
	translationText string
	translateCmd    = &cobra.Command{
		Use:   "translate",
		Short: "Translate word between languages",
		Long:  `You can also trasnalte phrases.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			srcLngCode, _ = cmd.Flags().GetString("source")
			dstLngCode, _ = cmd.Flags().GetString("dest")
			if !validateLangName(srcLngCode) {
				panic("Source language code is not valid!")
			}
			if !validateLangName(dstLngCode) {
				panic("Destination language code is not valid!")
			}

			fmt.Println(translateText(
				srcLngCode,
				dstLngCode,
				strings.Join(args, " "),
			))
		},
	}
)

// see https://github.com/GoogleCloudPlatform/golang-samples/blob/master/translate/v3/translate_text.go
// for more about translation sdk
func translateText(src string, dst string, text string) string {
	translatedText := ""
	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Fatal(err)
	}

	req := &translatepb.TranslateTextRequest{
		Parent: fmt.Sprintf(
			"projects/%s/locations/global",
			gcpProjectName),
		SourceLanguageCode: src,
		TargetLanguageCode: dst,
		MimeType:           "text/plain",
		Contents:           []string{text},
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		fmt.Errorf("TranslateText: %v", err)
	}

	// Display the translation for each input text provided
	for _, translation := range resp.GetTranslations() {
		translatedText += translation.GetTranslatedText()
	}

	return translatedText
}

func validateLangName(langCode string) bool {
	r, _ := regexp.Compile("^[a-zA-Z]+$")
	if r.MatchString(langCode) {
		return true
	}

	return false
}

func init() {
	translateCmd.Flags().StringVarP(
		&srcLngCode, "source", "s", "",
		"Source language as 2 - char code, e.g.: EN",
	)
	translateCmd.Flags().StringVarP(
		&dstLngCode, "dest", "d", "",
		"Destination language as 2 - char code, e.g.: EN",
	)
	rootCmd.MarkFlagRequired("source")
	rootCmd.MarkFlagRequired("dest")

	rootCmd.AddCommand(translateCmd)
}
