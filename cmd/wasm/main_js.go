//go:build js && wasm

package main

import (
 "bytes"
 "encoding/base64"
 "encoding/hex"
 "encoding/json"
 "fmt"
 "strings"
 "syscall/js"
 "unicode/utf8"

 "github.com/arsam92/-out-and-guss"
)

var extractFunc js.Func
func main() {
 extractFunc = js.FuncOf(extract)
 js.Global().Set("n17OutguessExtract", extractFunc)
 select {}
}
func extract(_ js.Value, args []js.Value) (ret any) {
 response := map[string]any{"ok": false}
 defer func() {
  if recovered := recover(); recovered != nil { response["error"] = fmt.Sprintf("Extractor stopped safely while parsing malformed input: %v", recovered) }
  encoded, err := json.Marshal(response)
  if err != nil { ret = "{\"ok\":false,\"error\":\"Could not serialize result.\"}"; return }
  ret = string(encoded)
 }()
 if len(args) < 1 || args[0].Type() != js.TypeObject { response["error"] = "Image bytes are missing."; return }
 view := args[0]
 length := view.Get("length").Int()
 if length <= 0 { response["error"] = "Image file is empty."; return }
 if length > 80*1024*1024 { response["error"] = "File exceeds the 80 MB safety limit."; return }
 input := make([]byte, length)
 if js.CopyBytesToGo(input, view) != length { response["error"] = "Could not read all image bytes."; return }
 var key []byte
 if len(args) > 1 && args[1].Type() == js.TypeString && args[1].String() != "" { key = []byte(args[1].String()) }
 var payload bytes.Buffer
 if err := outguess.Get(bytes.NewReader(input), &payload, key); err != nil { response["error"] = err.Error(); return }
 raw := payload.Bytes()
 response["ok"], response["size"], response["isText"] = true, len(raw), utf8.Valid(raw)
 response["text"], response["hex"], response["base64"] = strings.ToValidUTF8(string(raw), "�"), hex.EncodeToString(raw), base64.StdEncoding.EncodeToString(raw)
 return
}