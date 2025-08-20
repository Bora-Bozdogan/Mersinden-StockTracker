//auth import
import { auth } from "./index.js";

//listeners
export function InitDashboard(logoutFunc) {
    //get auth vars
    DisplayMerchant();
    DisplayItems();
    document.getElementById("log-button-js").addEventListener("click", AddItem);
    document.addEventListener("click", (e) => {
        //delete
        if (e.target.classList.contains("remove-button")) {
            DeleteItem(e);
        }

        //toggle update items
        if (e.target.classList.contains("update-button")) {
            //add editing class, update item 
            e.target.closest(".product").classList.add("editing")  
        }

        //cancel update items
        if (e.target.classList.contains("cancel-update-button")) {
            //cancel editing class
            e.target.closest(".product").classList.remove("editing")
        }

        //confirm update items
        if (e.target.classList.contains("confirm-update-button")) {
            //cancel editing class, update item 
            UpdateItem(e)
            e.target.closest(".product").classList.remove("editing")
        }

        //toggle update merchant
        if (e.target.classList.contains("update-merchant-button")) {
            //add editing class
            e.target.closest("#merchant-data").classList.add("editing")
        }
        
        //cancel update merchant
        if (e.target.classList.contains("cancel-update-merchant-button")) {
            //cancel editing class
            e.target.closest("#merchant-data").classList.remove("editing")
        }

        //confirm update merchant
        if (e.target.classList.contains("confirm-update-merchant-button")) {
            //cancel editing class, update item 
            UpdateMerchant(e)
            e.target.closest("#merchant-data").classList.remove("editing")
        }

        //logout
        if (e.target.closest('#logout-button')) {
            //call logout function in index
            logoutFunc()
        }
    })
}

const API_BASE = "http://127.0.0.1:3000";

//function to display all existing items 
async function DisplayItems () {
    //get user
    const user = auth.currentUser;
    const token = await user.getIdToken();

    let div = document.getElementById("results");
    div.innerHTML = "";

    try {
        const response = await fetch(API_BASE+"/items", {
            headers: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": `application/json`,
            }
        });
        if (!response.ok) {
            throw new Error(`HTTP ${response.status}`);
        }
        let productData = await response.json();
        for (let i = 0; i < productData.length; i++) {
            //parse into struct
            const {
                ID: id,
                ProductName: productName, 
                ProductDescription: productDescription, 
                MerchantID: merchantID, 
                Price: price, 
                Stock: stock
            } = productData[i]

            //get merchant values from id
            const response = await fetch(API_BASE+`/merchant/${merchantID}`, {
                headers: {
                "Authorization": `Bearer ${token}`,
                "Content-Type": 'application/json',
                }
            });
        
            if (!response.ok) {
                throw new Error(`HTTP ${response.status}`);
            }

            let merchantData = await response.json();
            const {
                MerchantName: merchantName,
                PhoneNumber: phoneNumber
            } = merchantData
        
        

            //add to div as html, switch between product view and edit
            div.innerHTML += 
            `
            <div class="product" data-id="${id}">
            <div class="view">
                <p>
                    <span class="label">Ürün:</span> <span class="value">${productName}</span>
                </p>
                <p>
                    <span class="label">Açıklama:</span> <span class="value">${productDescription}</span>
                </p>
                <p>
                    <span class="label">Fiyat:</span> <span class="value">${price}</span>
                </p>
                <p>
                    <span class="label">Stok:</span> <span class="value">${stock}</span>
                </p>
                <p>
                    <span class="label">Satıcı:</span> <span class="value">${merchantName}</span>
                </p>
                <p>
                    <span class="label">Telefon:</span> <span class="value">${phoneNumber}</span>
                </p>

                <button class="remove-button">Kaldır</button>
                <button class="update-button">Güncelle</button>
            </div>

            <div class="edit">
                <input type="text" class="product-name-value" placeholder="Urun" value="${productName}">
                <input type="text" class="product-description-value" placeholder="Urun Aciklamasi" value="${productDescription}">
                <input type="number" class="product-price-value" placeholder="Fiyat" value="${price}">
                <input type="number" class="product-stock-value" placeholder="Stok" value="${stock}">
                <p>
                    <span class="label">Satıcı:</span> <span class="value">${merchantName}</span>
                </p>
                <p>
                    <span class="label">Telefon:</span> <span class="value">${phoneNumber}</span>
                </p>

                <button class="confirm-update-button">Kaydet</button>
                <button class="cancel-update-button">İptal</button>
            </div>
        </div>

            `
        }    
        if (div.innerHTML == "") {
            div.innerHTML += 
            `
            <span> Stokta hiç ürün bulunmamaktadir </span> 
            `
        }
    } catch (err) {
        console.error('Fetch failed:', err);
    }
}

async function AddItem() {
    const user = auth.currentUser;
    const token = await user.getIdToken();

    //collect needed values 
    let productName = document.getElementById("product-text-js").value
    let productDescription = document.getElementById("product-description-text-js").value
    let price = parseInt(document.getElementById("price-text-js").value)
    let stock = parseInt(document.getElementById("stock-text-js").value)

    //reset textbox values
    document.getElementById("product-text-js").value = ""
    document.getElementById("product-description-text-js").value = ""
    document.getElementById("price-text-js").value = ""
    document.getElementById("stock-text-js").value = ""

    //send post request to go to add to db
    try {
        const response = await fetch(API_BASE+"/items", {
            method: 'POST',
            headers: {
                "Authorization": `Bearer ${token}`,
                'Content-Type': 'application/json'},
            body: JSON.stringify
            ({
                ProductName: productName,
                ProductDescription: productDescription,
                Price: price,
                Stock: stock
            })
        });    
        DisplayItems();
    } catch (err) {
        console.error('Fetch failed:', err);
    }
}

async function DeleteItem(e) {
    const user = auth.currentUser;
    const token = await user.getIdToken();

    //get relevant div id number
    let productDiv = e.target.closest(".product");
    const id = productDiv.dataset.id;

    //send delete request to backend with id
    try {
        const response = await fetch(API_BASE+`/items/${id}`, {
            method: 'DELETE',
            headers: {
                "Authorization": `Bearer ${token}`,
                'Content-Type': 'application/json'
            },
        });    
        DisplayItems();
    } catch (err) {
        console.error('Fetch failed:', err);
    }
}

async function UpdateItem(e) {
    const user = auth.currentUser;
    const token = await user.getIdToken();

    //get the id
    let productDiv = e.target.closest(".product");
    const id = productDiv.dataset.id;

    //get new values
    let productName = productDiv.querySelector(".product-name-value").value
    let productDescription = productDiv.querySelector(".product-description-value").value
    let price = parseInt(productDiv.querySelector(".product-price-value").value)
    let stock = parseInt(productDiv.querySelector(".product-stock-value").value)

    //send update request to id with new json
    try {
        const response = await fetch(API_BASE+`/items/${id}`, {
            method: 'PUT',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify
            ({
                ProductName: productName,
                ProductDescription: productDescription,
                Price: price,
                Stock: stock
            })
        });  
        DisplayItems();
    } catch (err) {
        console.error('Fetch failed:', err);
    }
}

//function to get merchant info and display it
async function DisplayMerchant() {
    const user = auth.currentUser;
    const token = await user.getIdToken();

    let div = document.getElementById("merchant-data");
    try {
        const response = await fetch(API_BASE+"/merchant/", 
        {   headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        }},
        );
        if (!response.ok) {
            throw new Error(`HTTP ${response.status}`);
        }
        let data = await response.json();
        //parse into struct
        const {
            MerchantName: merchantName, 
            PhoneNumber: phoneNumber, 
        } = data
        //add to div as html
        div.innerHTML = 
        `
        <div class="view">
            <p>
            Adiniz: ${merchantName} <br>
            Telefon numaraniz: ${phoneNumber}
            </p> <br>
            <button class="update-merchant-button">Güncelle</button>
        </div>

        <div class="edit">
            <h3>Kullanici bilgileri</h3> 
            <input type="text" id="merchant-name-js" placeholder="Adiniz:" value="${merchantName}">
            <input type="text" id="merchant-number-js" placeholder="Telefon numaraniz:" value="${phoneNumber}">
            <button class="confirm-update-merchant-button">Kaydet</button>
            <button class="cancel-update-merchant-button">Iptal</button>
        </div>
        `

    } catch (err) {
        console.error('Fetch failed:', err);
    }
}

async function UpdateMerchant(e) {
    const user = auth.currentUser;
    const token = await user.getIdToken();

    //get the id
    let merchantDiv = document.getElementById("merchant-data");

    //get new values
    let merchantName = merchantDiv.querySelector("#merchant-name-js").value
    let phoneNumber = merchantDiv.querySelector("#merchant-number-js").value

    //send update request to id with new json
    try {
        const response = await fetch(API_BASE+"/merchant", {
            method: 'PUT',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify
            ({
                MerchantName: merchantName,
                PhoneNumber: phoneNumber
            })
        });    
        DisplayMerchant();
    } catch (err) {
        console.error('Fetch failed:', err);
    }
}