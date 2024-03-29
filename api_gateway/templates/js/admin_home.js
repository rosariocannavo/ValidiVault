const searchBar = document.querySelector('.search-bar');

let account;
let token;

async function fetchData() {
    try {
        const response = await fetch('/get-cookie', {
            method: "GET",
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        account = data.account;
        token = data.token;

        console.log("Account:", account);
        console.log("Token:", token);
    } catch (error) {
        console.error("Error fetching data:", error);
    }
}

// Call the function when the page loads
window.addEventListener('load', () => {
    fetchData();
});


document.getElementById("setButton").addEventListener('click', async function () {
    document.getElementById('response').innerHTML = '';

    try {
        const productName = searchBar.value;
        searchBar.value = '';

        if (productName !== "") {
            document.getElementById('bar').style.border = '2px solid green';

            const url = `http://localhost:8080/admin/app/product?account=${account}&productName=${productName}`;

            const secondResponse = await fetch(url, {
                method: "PUT",
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `${token}`
                },
            });

            if (!secondResponse.ok) {
                throw new Error('Network response was not ok');
            }

            const responseData = await secondResponse.json();
            console.log(responseData);
            console.log("updated value")
            document.getElementById("productId").textContent = responseData.productId;
            document.getElementById("productName").textContent = responseData.productName;
            document.getElementById("manufacturer").textContent = responseData.manufacturer;
            document.getElementById("isRegistered").textContent = responseData.isRegistered;

            document.getElementById('response').innerHTML = '<p>Value updated on contract!</p>';

        } else {
            document.getElementById('bar').style.border = '2px solid red';

            document.getElementById('response').innerHTML = '<p>Invalid name</p>';
        }

    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
    }
});


document.getElementById("getButton").addEventListener('click', async function () {
    document.getElementById('response').innerHTML = '';
    document.getElementById('bar').style.border = '';

    try {
        const productId = parseInt(searchBar.value);
        searchBar.value = '';

        if (productId !== 0) {
            document.getElementById('bar').style.border = '2px solid green';

            const url = `http://localhost:8080/admin/app/product?productId=${productId}`;

            const secondResponse = await fetch(url, {
                method: "GET",
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `${token}`
                },
            });

            if (!secondResponse.ok) {
                throw new Error('Network response was not ok');
            }

            const responseData = await secondResponse.json();
            console.log(responseData);
            if (response.isRegistered == false) {
                document.getElementById("productId").textContent = "";
                document.getElementById("productName").textContent = "";
                document.getElementById("manufacturer").textContent = "";
                document.getElementById("isRegistered").textContent = "false";
                document.getElementById('response').innerHTML = '<p>Product not registered</p>';

            } else {
                document.getElementById("productId").textContent = responseData.productId;
                document.getElementById("productName").textContent = responseData.productName;
                document.getElementById("manufacturer").textContent = responseData.manufacturer;
                document.getElementById("isRegistered").textContent = responseData.isRegistered;
            }
        } else {
            document.getElementById('bar').style.border = '2px solid red';
            document.getElementById('response').innerHTML = '<p>Invalid id</p>';
        }
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
    }
});