// Copyright © 2018 Kristoffer Dalby <kradalby@kradalby.no>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"log"

	nb "github.com/kradalby/netboxxx/netbox"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write configuration based on template",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		host := viper.GetString("host")
		if host == "" {
			log.Fatalln("host is not set, please set the -n flag")
		}

		apikey := viper.GetString("apikey")
		if apikey == "" {
			log.Fatalln("apikey is not set, please set the -k flag")
		}

		c := nb.NewClient(host, apikey)

		c.WriteTemplates(TemplateFile)
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// writeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
