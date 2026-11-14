package starlarktest

import (
 "fmt"
 "os"
 "path/filepath"
 "strings"
 "testing"
 "go.starlark.net/starlark"
 "go.starlark.net/starlarkstruct"
)

type mockBrowser struct {
 scenario string
 page int
 url string
 ended, filled, submitted, filterOpen, filterApplied bool
 noForm, noList, noFilter bool
 calls []string
 results []starlark.Value
 checkpoints []int
}
func val(x any) starlark.Value {
 switch v:=x.(type) {
 case nil: return starlark.None
 case string: return starlark.String(v)
 case bool: return starlark.Bool(v)
 case int: return starlark.MakeInt(v)
 case []any:
  items:=make([]starlark.Value,0,len(v))
  for _,s:=range v {items=append(items,val(s))}
  return starlark.NewList(items)
 case map[string]any:
  d:=starlark.NewDict(len(v))
  for k,s:=range v {if err:=d.SetKey(starlark.String(k),val(s));err!=nil{panic(err)}}
  return d
 default: panic(fmt.Sprintf("invalid fixture type: %T",x))
 }
}
func obj(x map[string]any) starlark.Value {return val(x)}
func (m *mockBrowser) call(module,name string,args []starlark.Tuple) (starlark.Value,error) {
 m.calls=append(m.calls,module+"."+name)
 kw:=map[string]starlark.Value{}
 for _,pair:=range args {kw[string(pair[0].(starlark.String))]=pair[1]}
 str:=func(k string)string{if s,ok:=kw[k].(starlark.String);ok{return string(s)};return ""}
 if module=="job" {
  switch name{
  case "emit_result":
   if kw["data"]==nil{return nil,fmt.Errorf("data missing")}
   m.results=append(m.results,kw["data"])
   return obj(map[string]any{"accepted":true}),nil
  case "checkpoint":
   v,ok:=kw["page"].(starlark.Int)
   if !ok{return nil,fmt.Errorf("page must be int")}
   n,ok:=v.Int64()
   if !ok{return nil,fmt.Errorf("page overflow")}
   m.checkpoints=append(m.checkpoints,int(n))
   return obj(map[string]any{"page":int(n),"saved_at":"fixture"}),nil
  default:return nil,fmt.Errorf("unsupported job method %s",name)
  }
 }
 if name!="create_session"&&str("session_id")!="test-session" {
  return nil,fmt.Errorf("%s: wrong session_id",name)
 }
 switch name {
 case "create_session":
  m.page=1
  return obj(map[string]any{"session_id":"test-session","status":"running","default_tab_id":"main"}),nil
 case "navigate":
  m.url=str("url")
  if m.url==""{return nil,fmt.Errorf("URL required")}
  return obj(map[string]any{"url":m.url,"title":"Fixture","tab_id":"main","document_revision":1}),nil
 case "get_page_info":
  return obj(map[string]any{"url":m.url,"title":"Fixture title","tab_id":"main","document_revision":1}),nil
 case "end_session":
  m.ended=true
  return obj(map[string]any{"session_id":"test-session","status":"finished"}),nil
 case "get_links":
  return obj(map[string]any{"links":[]any{map[string]any{"ref":"link-1","text":"Example link","href":"/result","url":"https://example.com/result"}},"truncated":false}),nil
 case "inspect_lists":
  lists:=[]any{}
  if m.scenario=="list"&&!m.noList {
   lists=append(lists,map[string]any{"ref":fmt.Sprintf("list-%d",m.page),"text_sample":"Decretos municipais","item_count":2})
  }
  return obj(map[string]any{"lists":lists,"truncated":false}),nil
 case "get_list_items":
  if str("list_ref")!=fmt.Sprintf("list-%d",m.page){return nil,fmt.Errorf("stale list_ref")}
  return obj(map[string]any{"items":[]any{
   map[string]any{"ref":"item-1","title":"Decreto A","url":"https://example.com/a"},
   map[string]any{"ref":"item-2","title":"Decreto B","url":"https://example.com/b"},
  },"truncated":false}),nil
 case "inspect_pagination":
  return obj(map[string]any{"ref":fmt.Sprintf("page-%d",m.page),"has_next":m.page==1,"current_page":m.page,"mode":"link"}),nil
 case "next_page":
  if str("pagination_ref")!=fmt.Sprintf("page-%d",m.page){return nil,fmt.Errorf("stale pagination_ref")}
  m.page++
  return obj(map[string]any{"advanced":true,"url":m.url,"document_revision":m.page}),nil
 case "get_forms":
  forms:=[]any{}
  if m.scenario=="google"&&!m.noForm {
   forms=append(forms,map[string]any{"ref":"form-1","fields":[]any{map[string]any{"ref":"field-1","name":"q","tag":"textarea","type":"text"}}})
  }
  return obj(map[string]any{"forms":forms,"truncated":false}),nil
 case "fill_field":
  if str("element_ref")!="field-1"{return nil,fmt.Errorf("invalid element_ref")}
  m.filled=true
  return obj(map[string]any{"filled":true}),nil
 case "submit_form":
  if !m.filled||str("form_ref")!="form-1"{return nil,fmt.Errorf("invalid form submit")}
  m.submitted=true
  m.url="https://www.google.com/search?q=Starlark"
  return obj(map[string]any{"submitted":true,"navigated":true}),nil
 case "get_buttons":
  b:=[]any{}
  if m.scenario=="youtube"&&!m.noFilter {b=append(b,map[string]any{"ref":"filter-toggle","text":"Filtros","visible":true})}
  return obj(map[string]any{"buttons":b,"truncated":false}),nil
 case "get_elements":
  b:=[]any{}
  if m.filterOpen {b=append(b,map[string]any{"ref":"video-choice","text":"Vídeos","visible":true})}
  return obj(map[string]any{"elements":b,"truncated":false}),nil
 case "click":
  switch str("element_ref"){
  case "filter-toggle":m.filterOpen=true
  case "video-choice":
   if !m.filterOpen{return nil,fmt.Errorf("filter is not open")}
   m.filterApplied=true
  default:return nil,fmt.Errorf("unexpected element_ref")}
  return obj(map[string]any{"clicked":true,"navigated":m.filterApplied,"document_revision":2}),nil
 default:return nil,fmt.Errorf("browser method not mocked: %s",name)
 }
}
func (m *mockBrowser) globals() starlark.StringDict {
 names:=[]string{"create_session","navigate","get_page_info","end_session","get_links","inspect_lists",
  "get_list_items","inspect_pagination","next_page","get_forms","fill_field","submit_form","get_buttons","get_elements","click"}
 browser:=starlark.StringDict{}
 for _,name:=range names {
  name:=name
  browser[name]=starlark.NewBuiltin(name,func(_ *starlark.Thread,_ *starlark.Builtin,args starlark.Tuple,kwargs []starlark.Tuple)(starlark.Value,error){
   if len(args)!=0{return nil,fmt.Errorf("only keyword arguments allowed")}
   return m.call("browser",name,kwargs)
  })
 }
 job:=starlark.StringDict{}
 for _,name:=range []string{"emit_result","checkpoint"} {
  name:=name
  job[name]=starlark.NewBuiltin(name,func(_ *starlark.Thread,_ *starlark.Builtin,args starlark.Tuple,kwargs []starlark.Tuple)(starlark.Value,error){
   if len(args)!=0{return nil,fmt.Errorf("only keyword arguments allowed")}
   return m.call("job",name,kwargs)
  })
 }
 return starlark.StringDict{
  "browser":starlarkstruct.FromStringDict(starlarkstruct.Default,browser),
  "job":starlarkstruct.FromStringDict(starlarkstruct.Default,job),
 }
}
func (m *mockBrowser) run(file string) error {
 source,err:=os.ReadFile(filepath.Join("scripts",file))
 if err!=nil{return err}
 th:=&starlark.Thread{Name:"test-starlark"}
 th.SetMaxExecutionSteps(20000)
 _,err=starlark.ExecFile(th,file,source,m.globals())
 if e,ok:=err.(*starlark.EvalError);ok{return fmt.Errorf("%s",e.Backtrace())}
 return err
}
func field(t *testing.T,result starlark.Value,key string)string {
 t.Helper()
 d,ok:=result.(*starlark.Dict)
 if !ok{t.Fatalf("not a dict: %T",result)}
 x,found,err:=d.Get(starlark.String(key))
 if err!=nil||!found{t.Fatalf("missing field %q",key)}
 s,ok:=x.(starlark.String)
 if !ok{t.Fatalf("field %q not a string",key)}
 return string(s)
}
func called(calls []string,method string)bool {
 return strings.Contains("|"+strings.Join(calls,"|")+"|","|"+method+"|")
}
func TestStarlarkScenariosWithMocks(t *testing.T){
 cases:=[]struct{name,file,scenario string;noForm,noList,noFilter bool}{
  {"page_without_list","01-sem-lista.star","simple",false,false,false},
  {"paginated_list","02-lista-paginada.star","list",false,false,false},
  {"google_search","03-google-pesquisa.star","google",false,false,false},
  {"youtube_filter","04-youtube-filtro.star","youtube",false,false,false},
  {"google_without_form","03-google-pesquisa.star","google",true,false,false},
  {"missing_list","02-lista-paginada.star","list",false,true,false},
  {"youtube_without_filter","04-youtube-filtro.star","youtube",false,false,true},
 }
 for _,tc:=range cases{t.Run(tc.name,func(t *testing.T){
  m:=&mockBrowser{scenario:tc.scenario,noForm:tc.noForm,noList:tc.noList,noFilter:tc.noFilter}
  if err:=m.run(tc.file);err!=nil{t.Fatal(err)}
  if !m.ended||len(m.results)==0{t.Fatalf("execution missing results or cleanup: %+v",m)}
  for _,fn:=range []string{"browser.create_session","browser.navigate","job.emit_result","browser.end_session"}{
   if !called(m.calls,fn){t.Errorf("missing call %s",fn)}
  }
  switch tc.scenario {
  case "simple":
   if !called(m.calls,"browser.inspect_lists"){t.Error("list inspection missing")}
  case "list":
   if tc.noList{
    if field(t,m.results[0],"status")!="list_not_found"{t.Error("missing-list result not reported")}
   }else{
    if len(m.checkpoints)!=2||m.checkpoints[0]!=1||m.checkpoints[1]!=2{t.Errorf("checkpoints: %v",m.checkpoints)}
    if !called(m.calls,"browser.next_page"){t.Error("pagination missing")}
   }
  case "google":
   if tc.noForm {
    if field(t,m.results[0],"method")!="direct_url_fallback"{t.Error("fallback not reported")}
   }else if !m.filled||!m.submitted{t.Error("form not searched")}
  case "youtube":
   if tc.noFilter {
    if field(t,m.results[0],"status")!="filter_button_not_found"{t.Error("missing filter not reported")}
   }else if !m.filterOpen||!m.filterApplied{t.Error("video filter not clicked")}
  }
 })}
}
