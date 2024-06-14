/*  # # #   # # #  # # #  # # # #       # # # #     #       #        # # #   # # # #
  #     # #     # #    # #             #    # #   #        #       # * * #  #
 #       #     # #    # # # # #       # # #   # #         #       #  *  #  #  # # #
#     # #     # #    # #             #    #   #          #       # * * #  #      #
# # #   # # #  # # #  # # # #       # # #    #          # # # #  # # #    # # # # */

//ColorsByLOG.go --version 1.1.0

package main

const (
	DefaultColor string = "\x1b[0m"

	text_Black  string = "\x1b[30m"
	text_Red    string = "\x1b[31m"
	text_Green  string = "\x1b[32m"
	text_Yellow string = "\x1b[33m"
	text_Blue   string = "\x1b[34m"
	text_Purple string = "\x1b[35m"
	text_Cyan   string = "\x1b[36m"
	text_White  string = "\x1b[37m"

	attribute_Bold       string = "\x1b[1m"
	attribute_Italic     string = "\x1b[3m"
	attribute_Underlined string = "\x1b[4m"
	attribute_Invisible  string = "\x1b[8m"

	backgrond_Black  string = "\x1b[40m"
	backgrond_Red    string = "\x1b[41m"
	backgrond_Green  string = "\x1b[42m"
	backgrond_Yellow string = "\x1b[43m"
	backgrond_Blue   string = "\x1b[44m"
	backgrond_Purple string = "\x1b[45m"
	backgrond_Cyan   string = "\x1b[46m"
	backgrond_White  string = "\x1b[47m"
)

func SetColor(_colorName string) {
	print(_colorName)
}

func ResetColor() {
	print(DefaultColor)
}
