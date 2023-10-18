// Code generated by templ@v0.2.364 DO NOT EDIT.

package console

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Event(ev *LogEvent) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"flex flex-row text-blue-400 font-mono\" data-event-date=\"{ev.CreatedAt}\"><div class=\"w-1/6\">")
		if err != nil {
			return err
		}
		var var_2 string = ev.CreatedAt.Format("15:04:05.000")
		_, err = templBuffer.WriteString(templ.EscapeString(var_2))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div><div class=\"w-5/6 break-all\">")
		if err != nil {
			return err
		}
		var var_3 string = ev.EventData
		_, err = templBuffer.WriteString(templ.EscapeString(var_3))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func Console(currentRun int, runs []LogRun, events []LogEvent) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_4 := templ.GetChildren(ctx)
		if var_4 == nil {
			var_4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<!doctype html><html><head><title>")
		if err != nil {
			return err
		}
		var_5 := `gomon console`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><script defer src=\"https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js\">")
		if err != nil {
			return err
		}
		var_6 := ``
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><script src=\"https://cdn.tailwindcss.com\">")
		if err != nil {
			return err
		}
		var_7 := ``
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><link href=\"https://cdn.jsdelivr.net/npm/daisyui@3.9.3/dist/full.css\" rel=\"stylesheet\" type=\"text/css\"></head><body class=\"bg-slate-900 text-white flex flex-col h-screen\"><nav x-data class=\"grow-0 flex flex-row mx-2 p-2 justify-between items-center\"><div class=\"flex flex-row\"><a href=\"/\" class=\"text-2xl text-bold\">")
		if err != nil {
			return err
		}
		var_8 := `gomon`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></div><div class=\"flex flex-row items-center gap-2\"><label class=\"px-4\">")
		if err != nil {
			return err
		}
		var_9 := `Filter:`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><select class=\"select select-bordered text-slate-900\"><option value=\"all\" selected>")
		if err != nil {
			return err
		}
		var_10 := `all`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</option><option value=\"stdout\">")
		if err != nil {
			return err
		}
		var_11 := `stdout`
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</option><option value=\"stderr\">")
		if err != nil {
			return err
		}
		var_12 := `stderr`
		_, err = templBuffer.WriteString(var_12)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</option></select><label class=\"px-4\">")
		if err != nil {
			return err
		}
		var_13 := `Run:`
		_, err = templBuffer.WriteString(var_13)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><select class=\"select select-bordered text-slate-900\">")
		if err != nil {
			return err
		}
		for _, r := range runs {
			_, err = templBuffer.WriteString("<option value=\"{r.ID}\"")
			if err != nil {
				return err
			}
			if int(r.ID) == currentRun {
				if true {
					_, err = templBuffer.WriteString(" selected")
					if err != nil {
						return err
					}
				}
			}
			_, err = templBuffer.WriteString(">")
			if err != nil {
				return err
			}
			var var_14 string = r.CreatedAt.Format("2006-01-02 15:04:05")
			_, err = templBuffer.WriteString(templ.EscapeString(var_14))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</option>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</select><button id=\"btn-clear\" class=\"btn btn-primary\" @click=\"clearConsole\">")
		if err != nil {
			return err
		}
		var_15 := `Clear`
		_, err = templBuffer.WriteString(var_15)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button><button id=\"btn-restart\" class=\"btn btn-secondary\" @click=\"hardRestart\">")
		if err != nil {
			return err
		}
		var_16 := `Hard Restart`
		_, err = templBuffer.WriteString(var_16)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div></nav><main id=\"event-list\" class=\"grow flex flex-col justify-end mx-4 my-2 p-4 border-solid border border-blue-400 rounded-lg\">")
		if err != nil {
			return err
		}
		for _, ev := range events {
			err = Event(&ev).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</main><script>")
		if err != nil {
			return err
		}
		var_17 := `
				const eventList = document.getElementById("event-list");
				function listen() {
					const logSource = new EventSource("/__gomon__/events?stream=logs", {
						withCredentials: true,
					});

					const runSource = new EventSource("/__gomon__/events?stream=runs", {
						withCredentials: true,
					});

					logSource.onmessage = (event) => {
						eventList.insertAdjacentHTML("beforeend", event.data);
						eventList.scrollTop = eventList.scrollHeight;
					};

					runSource.onmessage = (event) => {
						window.location = window.location;
					};
				}

				function clearConsole() {
					eventList.innerHTML = "";
				}

				async function hardRestart() {
					console.log("hard restart");
					const res = await fetch("/restart", {
						method: "POST"
					})
				}

				document.addEventListener('alpine:init', () => {
					console.log("alpine init");
				});
			`
		_, err = templBuffer.WriteString(var_17)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script></body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}