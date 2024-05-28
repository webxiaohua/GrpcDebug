/**

 @Name：grpc_debug
 @Author：robin
 @Site：
 @License:

 */
layui.extend({
    selectInput: "selectInput/selectInput",
    jsonView:"jsonView/jsonView"
});
layui.define(['form'],function(exports){
    var $ = layui.$
        ,form = layui.form
        ,layer = layui.layer;
    var pathUserParamsMap={} // {pathStr:{user:params}} 基于path关联参数
    var pathList = []; // 用于渲染path
    var envIsLoading = false;

    // 处理env控件
    function handleNodeList(discoveryId){
        $.ajax({
            url:"/tool/grpc_debug/node_list",
            method:"POST",
            dataType:"json",
            contentType:"application/json",
            data:JSON.stringify({
                discovery_id:discoveryId,
            }),
            success:function (res) {
                if (res.code != 0) {
                    layer.msg(res.message);
                    $("#node_list").empty()
                }else{
                    if(res.data.path_list.length === 0) {
                        pathList.length = 0;
                    }else {
                        pathList.length = 0;
                        for (const v of res.data.path_list) {
                            pathStr = v.path;
                            pathList.push({value: pathStr, name: pathStr});
                            pathUserParamsMap[pathStr]=v.params;
                        }
                    }
                    initGrpcPath();
                    $("#node_list").empty()
                    for (let i in res.data.node_list) {
                        if (i == 0) {
                            $("#node_list").append("<input type=\"radio\" name=\"ip_addr\" value=\"" + res.data.node_list[i][1] + "\" title=\"" + res.data.node_list[i][0] + "\" lay-verify=\"required\" checked>");
                        } else {
                            $("#node_list").append("<input type=\"radio\" name=\"ip_addr\" value=\"" + res.data.node_list[i][1] + "\" title=\"" + res.data.node_list[i][0] + "\" lay-verify=\"required\">");
                        }
                    }
                }
                form.render('radio');
            }
        });
    }
    // 处理 discovery_id 控件
    function handleDiscoveryId(selectData) {
        layui.use(["selectInput"], function () {
            var selectInput = layui.selectInput;
            // 全量参数版本
            var ins = selectInput.render({
                // 容器id，必传参数
                elem: '#discovery_id',
                name: 'discovery_id', // 渲染的input的name值
                layFilter: 'select_discovery_id', //同layui form参数lay-filter
                layVerify: '', //同layui form参数lay-verify
                layVerType: 'tips', // 同layui form参数lay-verType
                layReqText: '请填写discovery_id', //同layui form参数lay-ReqText
                initValue: '', // 渲染初始化默认值
                hasSelectIcon: true,
                placeholder: '请输入discovery_id，按 Enter 自动索引环境列表', // 渲染的inputplaceholder值
                // 联想select的初始化本地值，数组格式，里面的对象属性必须为value，name，value是实际选中值，name是展示值，两者可以一样
                data: selectData,
                url: "", // 异步加载的url，异步加载联想词的数组值，设置url，data参数赋的值将会无效，url和data参数两者不要同时设置
                remoteSearch: false, // 是否启用远程搜索 默认是false，和远程搜索回调保存同步
                ignoreCase:true, // 是否忽略大小写
                parseData: function (data) {  // 此参数仅在异步加载data数据下或者远程搜索模式下有效，解析回调，如果你的异步返回的data不是上述的data格式，请在此回调格式成对应的数据格式，回调参数是异步加载的数据

                },
                error: function (error) { // 异步加载出错的回调 回调参数是错误msg

                },
                done: function (data) { // 异步加载成功后的回调 回调参数加载返回数据

                }
            });
            // 监听input 实时输入事件
            ins.on('itemInput(discovery_id)', function (obj) {
                //console.log(obj);
            });

            // 监听select 选择事件
            ins.on('itemSelect(discovery_id)', function (obj) {
                handleNodeList(obj.data);
            });
            $("#discovery_id").keydown(function (e){
                if(e.keyCode == 13) {
                    e.preventDefault(); // 阻止默认行为
                    console.log("enter");
                }
            })
            $("#discovery_id").keyup(function (e){
                if(e.keyCode == 13) {
                    handleNodeList(ins.getComponents().$inputElem.val());
                }
            })
            // 监听blur 光标离开事件
            /*ins.on('itemBlur(discovery_id)', function (obj) {
                console.log(2);

            });*/
        });
    }
// 处理 path 控件
    function initGrpcPath() {
        layui.use(["selectInput"], function () {
            var selectInput = layui.selectInput;
            // 全量参数版本
            renderOpts = {
                // 容器id，必传参数
                elem: '#grpc_path',
                name: 'grpc_path', // 渲染的input的name值
                layFilter: 'select_grpc_path', //同layui form参数lay-filter
                layVerify: 'required', //同layui form参数lay-verify
                layVerType: 'tips', // 同layui form参数lay-verType
                layReqText: '请填写grpc_path', //同layui form参数lay-ReqText
                initValue: '', // 渲染初始化默认值
                hasSelectIcon: true,
                placeholder: '请输入grpc_path', // 渲染的inputplaceholder值
                // 联想select的初始化本地值，数组格式，里面的对象属性必须为value，name，value是实际选中值，name是展示值，两者可以一样
                data: pathList,
                url: "", // 异步加载的url，异步加载联想词的数组值，设置url，data参数赋的值将会无效，url和data参数两者不要同时设置
                remoteSearch: false, // 是否启用远程搜索 默认是false，和远程搜索回调保存同步
                ignoreCase:true, // 是否忽略大小写
                parseData: function (data) {  // 此参数仅在异步加载data数据下或者远程搜索模式下有效，解析回调，如果你的异步返回的data不是上述的data格式，请在此回调格式成对应的数据格式，回调参数是异步加载的数据

                },
                error: function (error) { // 异步加载出错的回调 回调参数是错误msg

                },
                done: function (data) { // 异步加载成功后的回调 回调参数加载返回数据

                }
            }
            var ins = selectInput.render(renderOpts);
            // 监听select 选择事件
            ins.on('itemSelect(grpc_path)', function (obj) {
                $("#grpc_params").val(pathUserParamsMap[obj.data]);
            });
        });
    }
    // 初始化 discovery_id 控件
    $(function(){
        $.ajax({
            url:"/tool/grpc_debug/discovery_list",
            method:"POST",
            dataType:"json",
            contentType:"application/json",
            data:JSON.stringify({}),
            success:function (res) {
                var selectData = [];
                if (res.code != 0) {
                    layer.msg(res.message);
                }else{
                    for (let i in res.data) {
                        selectData.push({value:res.data[i],name:res.data[i]})
                    }
                }
                handleDiscoveryId(selectData);
            }
        });
    });
    // 提交
    form.on('submit(btn_run)',function (data) {
        console.log(data)
        //layer.msg(JSON.stringify(data.field));
        var btn = $(this)
        btn.text("处理中...").attr("disabled","disabled").addClass("layui-disabled");
        var grpcPath = data.field.grpc_path
        const regex = / - .+/i;
        grpcPath = grpcPath.replace(regex,'');
        $.ajax({
            url:"/tool/grpc_debug",
            method:"POST",
            dataType:"json",
            contentType:"application/json",
            xhrFields: {
                withCredentials: true
            },
            data:JSON.stringify({
                discovery_id:data.field.discovery_id,
                ip_addr:data.field.ip_addr_val,
                grpc_path:grpcPath,
                grpc_params:data.field.grpc_params
            }),
            success:function (res) {
                btn.text("运行").attr("disabled",false).removeClass("layui-disabled");
                if (res.code != 0) {
                    layer.msg(res.message);
                    $("#grpc_resp").text(res.message);
                    $("#grpc_resp_source").val(res.message);
                    $("#grpc_debug_path").val("grpcDebug --addr="+data.field.ip_addr_val+" --data='"+data.field.grpc_params+"' --path="+data.field.grpc_path);
                }else{
                    $("#grpc_resp").JSONView(res.data);
                    //const jsonObj = JSON.parse(res.data);
                    //$("#grpc_resp_source").val(JSON.stringify(jsonObj,null,2));
                    $("#grpc_resp_source").val(JSON.stringify(res.data,null,2));
                    $("#grpc_debug_path").val("grpcDebug --addr="+data.field.ip_addr_val+" --data='"+data.field.grpc_params+"' --path="+data.field.grpc_path);
                }
            }
        });
        return false;
    });

    initGrpcPath();

    // jsonView
    layui.use(["jsonView"],function (jsonView){
        $ = jsonView
    });

    // 复制内容
    /*$("#btn_copy").click(function (e){
        navigator.clipboard.writeText($("#grpc_resp_source").val()).then(() => {
            layer.msg('复制成功', { icon: 1, time: 1000 });
        });
    });*/
    //对外暴露的接口
    exports('grpc_debug', {});
});