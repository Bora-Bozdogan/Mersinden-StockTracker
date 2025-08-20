export function InitLogin(loginFunc, signupFunc) {
    document.getElementById("login-button").addEventListener("click", () => {
        GetCredentials(loginFunc)
    });
    document.getElementById("signup-button").addEventListener("click", () => {
        CreateAccount(signupFunc)
    });
}

function GetCredentials(loginFunc) {
    //get credentials
    const name = document.getElementById("login-name-text").value;
    const pass = document.getElementById("login-pass-text").value;

    //call login func
    loginFunc(name, pass)
}

function CreateAccount(signupFunc) {
    //get credentials
    const name = document.getElementById("login-name-text").value;
    const pass = document.getElementById("login-pass-text").value;

    //call login func
    signupFunc(name, pass)
}


