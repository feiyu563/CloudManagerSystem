//接入机型占比

var COLOR = {
	MACHINE:{
		TYPE_A:'#0175EE',
		TYPE_B:'#D89446',
		TYPE_C:'#373693',
		TYPE_D:'#25AE4F',
		TYPE_E:'#06B5C6',
		TYPE_F:'#009E9A',
		TYPE_G:'#AC266F'
	}
};

function renderLegend(){
	drawLegend(COLOR.MACHINE.TYPE_A,25,'Deployments');
	drawLegend(COLOR.MACHINE.TYPE_B,50,'Statefulset');
	drawLegend(COLOR.MACHINE.TYPE_C,75,'Daemonset');
	drawLegend(COLOR.MACHINE.TYPE_D,100,'Cronjob');
}

function drawLegend(pointColor,pointY,text){
	var ctx = $("#layer03_left_01 canvas").get(0).getContext("2d");
	ctx.beginPath();
	ctx.arc(20,pointY,6,0,2*Math.PI);
	ctx.fillStyle = pointColor;
	ctx.fill();
	ctx.font='20px';
	ctx.fillStyle = '#0c0c0c';
	ctx.fillText(text,40,pointY+3);
}


//系统资源
function renderLayer03Right(){
	drawLayer03Right($("#layer03_right_chart01 canvas").get(0),"#027825",0.66);
	drawLayer03Right($("#layer03_right_chart02 canvas").get(0),"#006DD6",0.52);
	drawLayer03Right($("#layer03_right_chart03 canvas").get(0),"#238681",0.34);
}

function drawLayer03Right(canvasObj,colorValue,rate){
	var ctx = canvasObj.getContext("2d");
    
	var circle = {
        x : 65,    //圆心的x轴坐标值
        y : 80,    //圆心的y轴坐标值
        r : 60      //圆的半径
    };


	ctx.beginPath();
	ctx.arc(circle.x,circle.y,circle.r,0,Math.PI*2)
	ctx.lineWidth = 10;
	ctx.strokeStyle = '#052639';
	ctx.stroke();
	ctx.closePath();

	ctx.beginPath();
	ctx.arc(circle.x,circle.y,circle.r,1.5*Math.PI,(1.5+rate*2)*Math.PI)
	ctx.lineWidth = 10;
	ctx.lineCap = 'round';
	ctx.strokeStyle = colorValue;
	ctx.stroke();
	ctx.closePath();
    
	ctx.fillStyle = 'blue';
	ctx.font = '20px Calibri';
	ctx.fillText(rate*100+'%',circle.x-15,circle.y+10);

}


function renderChartBar01(){
	var myChart = echarts.init(document.getElementById("layer03_left_02"));
		myChart.setOption(
					 {
						title : {
							text: '',
							subtext: '',
							x:'center'
						},
						tooltip : {
							trigger: 'item',
							formatter: "{b} : {c} ({d}%)"
						},
						legend: {
							show:false,
							x : 'center',
							y : 'bottom',
							data:['Deployments','Statefulset','Daemonset','Cronjob']
						},
						toolbox: {
						},
						label:{
							normal:{
								show: true, 
								formatter: "{b} \n{d}%"
							} 
						},
						calculable : true,
						color:[COLOR.MACHINE.TYPE_A,COLOR.MACHINE.TYPE_B,COLOR.MACHINE.TYPE_C,COLOR.MACHINE.TYPE_D],
						series : [
							{
								name:'',
								type:'pie',
								radius : [40, 80],
								center : ['50%', '50%'],
								//roseType : 'area',
								data:[
									{value:105, name:'Deployments'},
									{value:22, name:'Statefulset'},
									{value:4, name:'Daemonset'},
									{value:2, name:'Cronjob'},
								]
							}
						]
					}
		);

}


function renderLayer04Left(){
	var myChart = echarts.init(document.getElementById("layer04_left_chart"));
	myChart.setOption(
		{
			title: {
				text: ''
			},
			tooltip : {
				trigger: 'axis'
			},
			legend: {
				data:[]
			},
			grid: {
				left: '3%',
				right: '4%',
				bottom: '5%',
				top:'4%',
				containLabel: true
			},
			xAxis :
			{
				type : 'category',
				boundaryGap : false,
				data : getLatestDays(31),
				axisLabel:{
					textStyle:{
						color:"black", //刻度颜色
						fontSize:8  //刻度大小
					},
					rotate:45,
					interval:2
				},
				axisTick:{show:false},
				axisLine:{
					show:true,
					lineStyle:{
						color: '#0B3148',
						width: 1,
						type: 'solid'
					}
				}
			},
			yAxis : 
			{
				type : 'value',
				axisTick:{show:false},
				axisLabel:{
					textStyle:{
						color:"black", //刻度颜色
						fontSize:8  //刻度大小
						}
				},
				axisLine:{
					show:true,
					lineStyle:{
						color: '#0B3148',
						width: 1,
						type: 'solid'
					}
				},
				splitLine:{
					show:false
				}
			},
			tooltip:{
				formatter:'{c}',
				backgroundColor:'#FE8501'
			},
			series : [
				{
					name:'',
					type:'line',
					smooth:true,
					areaStyle:{
						normal:{
							color:new echarts.graphic.LinearGradient(0, 0, 0, 1, [{offset: 0, color: '#026B6F'}, {offset: 1, color: '#012138' }], false),
							opacity:0.2
						}
					},
					itemStyle : {  
                            normal : {  
                                  color:'#009991'
                            },
							lineStyle:{
								normal:{
								color:'#009895',
								opacity:1
							}
						}
                    },
					symbol:'none',
					data:[48, 52, 45, 46, 89, 120, 110,100,88,96,88,45,78,67,89,103,104,56,45,104,112,132,120,110,89,95,90,89,102,110,110]
				}
			]
		}
	
	);
}

function renderLayer04Right(){
	var myChart = echarts.init(document.getElementById("layer04_right_chart"));
	myChart.setOption({
			title: {
				text: ''
			},
			tooltip: {
				trigger: 'axis'
			},
			legend: {
				top:20,
				right:5,
				textStyle:{
					color:'black'
				},
				orient:'vertical',
				data:[
						{name:'Success',icon:'circle'},
						{name:'Failure',icon:'circle'},
					]
			},
			grid: {
				left: '3%',
				right: '16%',
				bottom: '3%',
				top:'3%',
				containLabel: true
			},
			xAxis: {
				type: 'category',
				boundaryGap: false,
				axisTick:{show:false},
				axisLabel:{
					textStyle:{
						color:"black", //刻度颜色
						fontSize:8  //刻度大小
						}
				},
				axisLine:{
					show:true,
					lineStyle:{
						color: '#0B3148',
						width: 1,
						type: 'solid'
					}
				},
				data: get10MinutesScale()
			},
			yAxis: {
				type: 'value',
				axisTick:{show:false},
				axisLabel:{
					textStyle:{
						color:"black", //刻度颜色
						fontSize:8  //刻度大小
						}
				},
				axisLine:{
					show:true,
					lineStyle:{
						color: '#0B3148',
						width: 1,
						type: 'solid'
					}
				},
				splitLine:{
					show:false
				}
			},
			series: [
						{
							name:'Success',
							type:'line',
							itemStyle : {  
									normal : {  
									color:'#F3891B'
								},
								lineStyle:{
									normal:{
									color:'#F3891B',
									opacity:1
										}
								}
							},  
							data:[5898, 4898, 3500, 4058, 2856, 1711, 5560]
						},
						{
							name:'Failure',
							type:'line',
							itemStyle : {  
									normal : {  
									color:'#006AD4'
								},
								lineStyle:{
									normal:{
									color:'#F3891B',
									opacity:1
										}
								}
							},
							data:[220, 182, 191, 234, 290, 10, 310]
						},
					]
		}	
	);
}

function get10MinutesScale()
{
	var currDate = new Date();
	var odd = currDate.getMinutes()%10;
	var returnArr = new Array();
	currDate.setMinutes(currDate.getMinutes()-odd);
	for(var i = 0; i <7; i++){
		returnArr.push(currDate.getHours()+":"+(currDate.getMinutes()<10?("0"+currDate.getMinutes()):currDate.getMinutes()));
		currDate.setMinutes(currDate.getMinutes()-10);
	}
	return returnArr;
}


function getLatestDays(num)
{
	var currentDay = new Date();
	var returnDays = [];
	for (var i = 0 ; i < num ; i++)
	{
		currentDay.setDate(currentDay.getDate() - 1);
		returnDays.push((currentDay.getMonth()+1)+"/"+currentDay.getDate());
	}
	return returnDays;
}