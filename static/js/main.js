function chartCounts(source, target) {
	for (var k in source) {
		console.log({"label": k, "value": source[k]})
		target.push({"label": k, "value": source[k]})
	}


	return MG.data_graphic({
	    // title: "Bar Prototype",
	    // description: "Work-in-progress",
	    data: target,
	    chart_type: 'bar',
	    x_accessor: 'value',
	    y_accessor: 'label',
	    // baseline_accessor: 'baseline',
	    // predictor_accessor: 'prediction',
	    width: 300,
	    right: 50,
	    target: document.getElementById('metrics'),
	    animate_on_load: true,
	    x_axis: false
	})
}

var editor = $("textarea#code")[0],
	metric_outcomes = [],
	config = {
		lineNumbers: true,
		keyMap: "sublime",
		mode: "text/x-sql",
		theme: "ambiance",
		indentWithTabs: false
	}

editor.value = "select vp.desc as aa_agency, au.auth_no\nfrom au_master au, vp_agency_master vp\nwhere\n	au.agency_id = vp.agency_id and\n	year(start_date) = 2015"
var cm = CodeMirror.fromTextArea(editor, config);

var keysDown = {}

cm.on("keydown", function (cm, event) {
	keysDown[event.keyCode] = true
	if (keysDown[13] && keysDown[16]) {
		event.preventDefault()
		$("form#editor button").click()
	}
})

cm.on("keyup", function (cm, event) {
	keysDown[event.keyCode] = false
})


$("form#editor").on("submit", scriptSubmit)

function scriptSubmit(event) {
	event.preventDefault();
	var $form = $( this ),
		code = editor.value
		lang = $form.find( "select[name='lang']" ).val(),
		payload = { code: code, lang: lang },
		metric_counter = {}


			$.post( "/", payload, function( data ) {
				data = JSON.parse(data)
				data["output"] = JSON.parse(data["output"])

				var output = data["output"]
				var tbl = ""
				var got_headers = false
				var headers = ""
				var limit = $("input#limit").val()


				for (var row in output) {
					if (limit == 0) { break }
					if (!got_headers) {
						tbl += "<thead><tr>"					
						for (var head in output[row]) {
							tbl += "<th>" + head + "</th>"
						}
						tbl += "</tr><thead><tbody>"					
						got_headers = true
					}
					tbl += "<tr>"

					var first_column = true				
					for (var field in output[row]) {
						tbl += "<td>" + output[row][field] + "</td>"
						if (first_column) {
							if (output[row][field] in metric_counter) {
								metric_counter[output[row][field]] += 1
							} else {
								metric_counter[output[row][field]] = 1
							}
							first_column = false
						}
					}
					tbl += "</tr>"
					limit -= 1
				}
				$("table#output").html(tbl+"</tbody>")
				// console.log(metric_counter)
				return

			}).done(function () {
			chartCounts(metric_counter, [])
			} )
	// $("form#editor").clone().appendTo(".container")

	


	// d3.json('static/data/fake_users1.json', function(fake) {
	// 	console.log(fake)
//  			fake = MG.convert.date(fake, 'date');
	//     MG.data_graphic({
	//         title: "Line Chart",
	//         description: "This is a simple line chart. You can remove the area portion by adding area: false to the arguments list.",
	//         data: fake,
	//         width: 800, // ADD PERCENT FUNCTIONALITY
	//         height: 500,
	//         right: 40,
	//         target: document.getElementById('metrics'),
	//         x_accessor: 'date',
	//         y_accessor: 'value'
	//     });
	// });




}


