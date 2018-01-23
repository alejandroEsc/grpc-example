
// Do knock a door
console.log("sending message: {knockDoor:true}")
client.getHello({knockDoor:true}, printReply);

// Do NOT knock a door
console.log("sending message: {knockDoor:false}")
client.getHello({knockDoor:false}, printReply);
