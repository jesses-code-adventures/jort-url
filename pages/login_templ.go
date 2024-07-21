// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Login() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script type=\"text/javascript\">\n            let loginMode = false;\n\n            function setLoginMode() {\n                loginMode = true;\n                setVisualStates();\n            }\n\n            function setSignUpMode() {\n                loginMode = false;\n                setVisualStates();\n            }\n\n            function setVisualStates() {\n                updateButtonText();\n                updateFormAction();\n                updateButtonStyles();\n            }\n\n            function getButtonText() {\n                if (loginMode) {\n                    return \"login\";\n                } else {\n                    return \"sign up\";\n                }\n            }\n\n            function updateButtonText() {\n                document.getElementById(\"submitButton\").innerText = getButtonText();\n            }\n\n            function updateFormAction() {\n                const form = document.getElementById(\"authForm\");\n                if (loginMode) {\n                    form.action = \"/login\";\n                } else {\n                    form.action = \"/user\";\n                }\n            }\n\n            function updateButtonStyles() {\n                const loginButton = document.getElementById(\"loginSelector\");\n                const signUpButton = document.getElementById(\"signUpSelector\");\n\n                if (loginMode) {\n                    loginButton.classList.add(\"bg-violet-200\");\n                    loginButton.classList.add(\"dark:bg-violet-800\");\n                    signUpButton.classList.remove(\"bg-violet-200\");\n                    signUpButton.classList.remove(\"dark:bg-violet-800\");\n                } else {\n                    signUpButton.classList.add(\"dark:bg-violet-800\");\n                    signUpButton.classList.add(\"bg-violet-200\");\n                    loginButton.classList.remove(\"dark:bg-violet-800\");\n                    loginButton.classList.remove(\"bg-violet-200\");\n                }\n            }\n\n            function handleFormSubmit(event) {\n                event.preventDefault();\n                const form = event.target;\n                const formData = new FormData(form);\n                const action = form.action;\n                const errorContainer = document.getElementById(\"errorContainer\");\n\n                fetch(action, {\n                    method: \"POST\",\n                    body: formData,\n                })\n                .then(response => {\n                    if (!response.ok) {\n                        return response.text().then(text => { throw new Error(text); });\n                    }\n                    return response.text();\n                })\n                .then(data => {\n                    window.location.href = \"/\";\n                })\n                .catch(error => {\n                    if (errorContainer === null) {\n                        console.error(error);\n                        return;\n                    }\n                    if (errorContainer.classList.contains(\"hidden\")) {\n                        errorContainer.classList.remove(\"hidden\");\n                    }\n                    errorContainer.innerText = error.message;\n                    errorContainer.style.display = \"block\";\n                });\n            }\n\n            document.addEventListener(\"DOMContentLoaded\", function() {\n                updateButtonText();\n                updateButtonStyles();\n                document.getElementById(\"authForm\").addEventListener(\"submit\", handleFormSubmit);\n            });\n        </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Layout("login - jort url").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"loginContainerContainer\" class=\"flex justify-center\"><div id=\"loginContainer\" class=\"flex w-full flex-col rounded-md lg:w-3/5 lg:border-2 lg:border-black\"><div id=\"loginOrSignUpSelectorContainer\" class=\"flex justify-evenly border-b-2 border-b-black\"><button id=\"signUpSelector\" onclick=\"setSignUpMode()\" class=\"w-full bg-violet-200 p-2 text-center text-xl dark:bg-violet-800\">sign up</button> <button id=\"loginSelector\" onclick=\"setLoginMode()\" class=\"w-full p-2 text-center text-xl\">log in</button></div><form id=\"authForm\" action=\"/user\" method=\"post\" class=\"pt-4\"><div id=\"usernameInputContainer\" class=\"flex flex-col items-center gap-2\"><label for=\"username\">username</label> <input type=\"text\" name=\"username\" id=\"username\" autocomplete=\"username\" class=\"w-1/3 rounded-md border-2 p-1 focus:outline-violet-200 dark:focus:outline-violet-800 dark:bg-black dark:text-white\" required></div><div id=\"passwordInputContainer\" class=\"flex flex-col items-center gap-2 pb-4\"><label for=\"password\">password</label> <input type=\"password\" name=\"password\" id=\"password\" autocomplete=\"current-password\" class=\"w-1/3 rounded-md border-2 p-1 focus:outline-violet-200 dark:focus:outline-violet-800 dark:bg-black dark:text-white\" required></div><button type=\"submit\" id=\"submitButton\" class=\"w-32 rounded-md bg-violet-500 px-4 py-2 text-white hover:bg-violet-300 dark:hover:bg-violet-700\">loading...</button></form><div id=\"errorContainer\" class=\"hidden text-center text-red-500\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
