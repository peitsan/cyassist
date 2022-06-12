export const formatDate=(timeStamp)=>{
	let Date= JSON.stringify(timeStamp);
	let a= Date.substring(0,9);
	let b= Date.substring(11,15);
   return a+' '+b
}
