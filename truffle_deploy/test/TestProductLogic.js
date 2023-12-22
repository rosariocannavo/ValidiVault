const ProductLogic = artifacts.require('ProductLogic');
//truffle test
contract('ProductLogic', (accounts) => {
    let productLogicInstance;

    before(async () => {
        productLogicInstance = await ProductLogic.deployed();
    });

    it('should register a product', async () => {
        const productId = 1;
        const productName = 'Example Product';

        // Register a product
        await productLogicInstance.registerProduct(productId, productName, { from: accounts[0] });

        // Retrieve the registered product
        const retrievedProduct = await productLogicInstance.getProduct(productId);

        assert.equal(retrievedProduct.productName, productName, 'Product name should match');
        assert.equal(retrievedProduct.isRegistered, true, 'Product should be registered');
    });

    it('should prevent registering an already registered product', async () => {
        const productId = 2;
        const productName = 'Registered Product';
    
        // Register a product first
        await productLogicInstance.registerProduct(productId, productName, { from: accounts[0] });
    
        // Try registering the same product again and expect it to revert
        try {
            await productLogicInstance.registerProduct(productId, productName, { from: accounts[0] });
            assert.fail('Expected revert not received');
        } catch (error) {
            assert(error.message.includes('Product already registered'), `${error.message}`);
        }
    });

    it('should retrieve a registered product', async () => {
        const productId = 1;
        const productName = 'Example Product';

        // Retrieve the registered product
        const retrievedProduct = await productLogicInstance.getProduct(productId);

        assert.equal(retrievedProduct.productName, productName, 'Product name should match');
        assert.equal(retrievedProduct.isRegistered, true, 'Product should be registered');
    });

    it('should toggle emergency stop', async () => {
        const initialStopStatus = await productLogicInstance.stopped();

        // Toggle emergency stop
        await productLogicInstance.toggleContractStopped({ from: accounts[0] });

        const updatedStopStatus = await productLogicInstance.stopped();
        assert.notEqual(initialStopStatus, updatedStopStatus, 'Stop status should change');
    });
});

