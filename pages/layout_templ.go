// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Layout(page string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
		templ_7745c5c3_Err = Globals(NewGlobalsArgs(page)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script type=\"text/javascript\">\n        function getCookie(cname) {\n          let name = cname + \"=\";\n          let decodedCookie = decodeURIComponent(document.cookie);\n          let ca = decodedCookie.split(';');\n          for(let i = 0; i <ca.length; i++) {\n            let c = ca[i];\n            while (c.charAt(0) == ' ') {\n              c = c.substring(1);\n            }\n            if (c.indexOf(name) == 0) {\n              return c.substring(name.length, c.length);\n            }\n          }\n          return \"\";\n        }\n        let loggedIn = getCookie(\"jort_user_id\") !== \"\";\n        document.addEventListener(\"DOMContentLoaded\", function() {\n            const hamburger = document.getElementById(\"hamburger\");\n            if (hamburger === null) {\n                throw new Error(\"Hamburger not found\");\n            }\n            hamburger.addEventListener(\"click\", function() {\n                const dropdownMenu = document.getElementById(\"dropdownContainer\");\n                if (dropdownMenu === null) {\n                    throw new Error(\"Dropdown menu not found\");\n                }\n                dropdownMenu.classList.toggle(\"hidden\");\n            });\n            setShowHamburger();\n        })\n        document.addEventListener(\"reload\", function() {\n            setShowHamburger();\n        })\n        function setShowHamburger() {\n            console.log('setShowHamburger');\n            if (loggedIn) {\n                console.log(\"logged in\");\n                document.getElementById(\"hamburger\").hidden = false;\n                return\n            } else {\n                console.log(\"not logged in\");\n                document.getElementById(\"hamburger\").hidden = true;\n                return\n            }\n        }\n        function toggleDarkMode() {\n            document.body.classList.toggle(\"dark\");\n            const dropdownMenu = document.getElementById(\"dropdownContainer\");\n            if (dropdownMenu === null) {\n                throw new Error(\"Dropdown menu not found\");\n            }\n            if (!dropdownMenu.classList.contains(\"hidden\")) {\n                dropdownMenu.classList.toggle(\"hidden\");\n            }\n        }\n        function handleLogout() {\n            fetch(\"/logout\", {\n                method: \"POST\",\n            }).then((response) => {\n                if (response.ok) {\n                    window.location.href = \"/login\";\n                }\n            });\n        }\n    </script><div id=\"headerContainer\" class=\"relative flex h-12 justify-end\"><img id=\"hamburger\" src=\"/static/hamburger.svg\" class=\"pr-12 hover:cursor-pointer\"><div id=\"dropdownContainer\" class=\"absolute top-full hidden pr-12\"><div id=\"dropdownMenu\" class=\"flex w-32 flex-col rounded-md border-2 border-black\"><button onclick=\"toggleDarkMode()\" class=\"px-4 py-1 hover:bg-black hover:text-white\">Dark Mode</button> <button onclick=\"handleLogout()\" class=\"px-4 py-1 hover:bg-black hover:text-white\">Logout</button></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
