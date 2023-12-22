const express = require('express');

const { web3, contractAddress, contractABI } = require('./ethConfig');

const app = express();
const port = 3000;

// Set up the contract instance
const contract = new web3.eth.Contract(contractABI, contractAddress);

// Endpoint to interact with the contract, this route will be contacted by the proxy 
app.get('/registerProduct', async (req, res) => {
    try {
        const account = req.query.account; // Get account from query parameters or request body
		console.log("account:" + account)
		
		const productName = req.query.productName
		console.log("productName: " + productName);

		//the product id is created by the first 10 char of the hash of account address and product name
		const productId = web3.utils.soliditySha3(account, productName).substring(0, 10); 
		console.log("productId:" + productId)

        // Set a new value using the registerProduct function (transaction)
        const tx = await contract.methods.registerProduct(productId, productName).send({
            from: account,
            gas: 200000,
        });

        // Check the updated value after the transaction
        const updatedValue = await contract.methods.getProduct(productId).call();
		
        console.log('Updated value:', updatedValue);

        res.status(200).json({  success: true, 
				"productId": updatedValue.productId.toString(), 
				"productName": updatedValue.productName, 
				"manufacturer": updatedValue.manufacturer, 
				"isRegistered": updatedValue.isRegistered });
    } catch (error) {
        console.error('Error:', error);
        res.status(500).json({ success: false, error: error.message });
    }
});

app.get('/getProduct', async (req, res) => {
    try {
		//const account = req.query.account; // Get account from query parameters or request body
		const productId = req.query.productId
		console.log("productId" + productId);

		const retrievedValue = await contract.methods.getProduct(productId).call();

		console.log('Retrieved value:', retrievedValue);

		res.status(200).json({  success: true, 
			"productId": retrievedValue.productId.toString(), 
			"productName": retrievedValue.productName, 
			"manufacturer": retrievedValue.manufacturer, 
			"isRegistered": retrievedValue.isRegistered });

	} catch (error) {
        console.error('Error:', error);
        res.status(200).json({ success: false, "isRegistered" : false });
    }

});


app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});

