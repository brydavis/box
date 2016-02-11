function tablize (source) {
	var target = []
	for (var k in source) {
		target.push({"label": k, "value": source[k]})
	}
	return target
}


function head (obj, lim) {
	return obj.sort(function (a, b) {
	  return b["value"] - a["value"];
	}).slice(0,lim)
}

function tail (obj, lim) {
	return obj.sort(function (a, b) {
	  return a["value"] - b["value"];
	}).slice(0,lim)
}



