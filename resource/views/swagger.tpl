<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>({{.Env}})API docs for {{.AppName}} - By Swagger UI</title>
  <link href="https://fonts.cat.net/css?family=Open+Sans:400,700|Source+Code+Pro:300,600|Titillium+Web:400,600,700" rel="stylesheet">
  <link rel="stylesheet" type="text/css" href="{{.SwgUIPath}}/swagger-ui.css" >
  <link rel="icon" type="image/png" href="{{.SwgUIPath}}/favicon-32x32.png" sizes="32x32" />
  <link rel="icon" type="image/png" href="{{.SwgUIPath}}/favicon-16x16.png" sizes="16x16" />
  <style>
    html
    {
      box-sizing: border-box;
      overflow: -moz-scrollbars-vertical;
      overflow-y: scroll;
    }
    *,
    *:before,
    *:after
    {
      box-sizing: inherit;
    }

    body {
      margin:0;
      background: #fafafa;
      font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol"
    }
    code {
      font-size: 90%;
      padding: 0.2em 0.4em;
      margin: 0;
      color: #c7254e;
      background-color: #f9f5f4;
      border-radius: 3px;
      font-family: Menlo,Monaco,Consolas,"Courier New",monospace;
    }
    #goto-top {
      display: inline-block;
      font-size: larger;
      font-weight: 800;
      border: 1px solid #ddd;
      border-radius: 5px;
      position: fixed;
      color: #666;
      right: 5px;
      bottom: 50px;
      padding: 2px 6px;
      text-decoration: none;
      background-color: #afadad80;
      z-index: 100;
    }
    .custom-info {
      padding: 15px 0;
    }
  </style>
</head>

<body>

<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" style="position:absolute;width:0;height:0">
  <defs>
    <symbol viewBox="0 0 20 20" id="unlocked">
          <path d="M15.8 8H14V5.6C14 2.703 12.665 1 10 1 7.334 1 6 2.703 6 5.6V6h2v-.801C8 3.754 8.797 3 10 3c1.203 0 2 .754 2 2.199V8H4c-.553 0-1 .646-1 1.199V17c0 .549.428 1.139.951 1.307l1.197.387C5.672 18.861 6.55 19 7.1 19h5.8c.549 0 1.428-.139 1.951-.307l1.196-.387c.524-.167.953-.757.953-1.306V9.199C17 8.646 16.352 8 15.8 8z"></path>
    </symbol>

    <symbol viewBox="0 0 20 20" id="locked">
      <path d="M15.8 8H14V5.6C14 2.703 12.665 1 10 1 7.334 1 6 2.703 6 5.6V8H4c-.553 0-1 .646-1 1.199V17c0 .549.428 1.139.951 1.307l1.197.387C5.672 18.861 6.55 19 7.1 19h5.8c.549 0 1.428-.139 1.951-.307l1.196-.387c.524-.167.953-.757.953-1.306V9.199C17 8.646 16.352 8 15.8 8zM12 8H8V5.199C8 3.754 8.797 3 10 3c1.203 0 2 .754 2 2.199V8z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="close">
      <path d="M14.348 14.849c-.469.469-1.229.469-1.697 0L10 11.819l-2.651 3.029c-.469.469-1.229.469-1.697 0-.469-.469-.469-1.229 0-1.697l2.758-3.15-2.759-3.152c-.469-.469-.469-1.228 0-1.697.469-.469 1.228-.469 1.697 0L10 8.183l2.651-3.031c.469-.469 1.228-.469 1.697 0 .469.469.469 1.229 0 1.697l-2.758 3.152 2.758 3.15c.469.469.469 1.229 0 1.698z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="large-arrow">
      <path d="M13.25 10L6.109 2.58c-.268-.27-.268-.707 0-.979.268-.27.701-.27.969 0l7.83 7.908c.268.271.268.709 0 .979l-7.83 7.908c-.268.271-.701.27-.969 0-.268-.269-.268-.707 0-.979L13.25 10z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="large-arrow-down">
      <path d="M17.418 6.109c.272-.268.709-.268.979 0s.271.701 0 .969l-7.908 7.83c-.27.268-.707.268-.979 0l-7.908-7.83c-.27-.268-.27-.701 0-.969.271-.268.709-.268.979 0L10 13.25l7.418-7.141z"/>
    </symbol>


    <symbol viewBox="0 0 24 24" id="jump-to">
      <path d="M19 7v4H5.83l3.58-3.59L8 6l-6 6 6 6 1.41-1.41L5.83 13H21V7z"/>
    </symbol>

    <symbol viewBox="0 0 24 24" id="expand">
      <path d="M10 18h4v-2h-4v2zM3 6v2h18V6H3zm3 7h12v-2H6v2z"/>
    </symbol>

  </defs>
</svg>

<div id="swagger-ui"></div>

<a id="goto-top" href="#"> &and; </a>

<script src="{{.SwgUIPath}}/swagger-ui-bundle.js"> </script>
<script src="{{.SwgUIPath}}/swagger-ui-standalone-preset.js"> </script>
<script src="{{.AssetPath}}/libs/jquery-3.3.1.slim.min.js"> </script>
<script>
window.onload = function() {

  // Build a system
  window.ui = SwaggerUIBundle({
    // url: "http://petstore.swagger.io/v2/swagger.json",
    url: "{{.JsonFile}}",
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout",
    // layout: "BaseLayout",

    /**
     * user custom config. more see @link https://swagger.io/docs/swagger-tools/#usage-34
     */
    // disable validate
    validatorUrl: null,
    displayRequestDuration: true,
    showRequestHeaders: true,
    jsonEditor: true
  })

  setTimeout(on_swagger_rendered, 2000)
}

function on_swagger_rendered() {
  let count = $('div.opblock').length
  let countTag = $('h4.opblock-tag').length
  let countGet = $('div.opblock-summary-get').length
  let countPost = $('div.opblock-summary-post').length
  let countPut = $('div.opblock-summary-put').length
  let countDel = $('div.opblock-summary-delete').length
  let modelNum = $('div.model-container').length
  let upTime = '{{.UpdateTime}}'

  let html = `<div class="custom-info">
<h2>接口信息</h2>
<ul>
    <li><strong>更新时间: </strong> <code>${upTime}</code></li>
    <li><strong>接口环境: </strong> <code>{{.Env}}</code></li>
    <li><strong>接口统计: </strong> <code>${count}</code> 个
        <small>(<span>GET: </span> <code>${countGet}</code>，
        <span>POST: </span> <code>${countPost}</code>，
        <span>PUT: </span> <code>${countPut}</code>，
        <span>DELETE: </span> <code>${countDel}</code>)</small>
    </li>
    <li><strong>模型统计: </strong> <code>${modelNum}</code></li>
    <li><strong>标签统计: </strong> <code>${countTag}</code></li>
</ul>
</div>
`

  $('.information-container').append(html)
}
</script>
</body>

</html>
