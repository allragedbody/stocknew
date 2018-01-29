<!DOCTYPE html>
<html style="height: 100%">
   <head>
       <meta charset="utf-8">
   </head>
<body>
<form id="stock" action="/draw/000158.SZ/60">
    名称：<input name="code" type="text" />
	日期：<input name="date" type="date" value="2018-01-01" />
    时间段：
	<select>
  		<option value="day60">60</option>
  		<option value="day120">120</option>
  		<option value="day180">180</option>
  		<option value="day360">360</option>
	</select>
    <input type="submit" value="提交" />
</form>
</body>
</html>