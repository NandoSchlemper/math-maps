// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.906
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "math-coordenadas/frontend/components"

func Initial() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"pt-BR\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>THOR</title><link rel=\"icon\" href=\"./static/icon.png\" type=\"image/png\"></head><body><div class=\"centered\"><div class=\"message-div\"><p class=\"message-content\">Bem vindo à  <span style=\"color:blueviolet; font-weight:bold;\">Thor</span></p></div></div><div class=\"bottom\"><div class=\"button\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.InitialButton().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div></div></body></html><style lang=\"css\">\r\n        html, body {\r\n            margin: 0;\r\n            padding: 0;\r\n            height: 100%;\r\n            width: 100%; \r\n        }\r\n        p {\r\n            font-family: sans-serif;\r\n            text-align: center;\r\n            font-size: medium;\r\n        }\r\n        .centered {\r\n            position: relative;\r\n            display: flex;\r\n            justify-content: center;\r\n            align-items: center;\r\n            height: 80vh;\r\n        }\r\n        .bottom {\r\n            height: 15vh;\r\n        }\r\n\r\n        .message-content {\r\n            font-size: 5vh;\r\n        }\r\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
