//imports
//firebase
import { initializeApp } from 'firebase/app';
import { 
    getAuth, 
    onAuthStateChanged,
    signInWithEmailAndPassword,
    connectAuthEmulator,
    createUserWithEmailAndPassword,
    signOut
} from 'firebase/auth'; //auth modules
import { InitDashboard } from '../pages/dashboard'; //dashboard
import { InitLogin } from '../pages/login'; //dashboard

//initialize firebase app here
const firebaseApp = initializeApp ({
});

//initialize auth and //connect to auth emulator for testing
export const auth = getAuth(firebaseApp);
connectAuthEmulator(auth, "http://127.0.0.1:9099");

//dynamic auth change function listener
onAuthStateChanged(auth, checkLogin)

async function checkLogin(user) {
    let appDiv = document.getElementById("app");
    
    if (user != null) {
        //logged in, fetch dashboard html, update index.html inside 
        const res = await fetch("/templates/dashboard.html");
        if (!res.ok) {
            throw new Error(`HTTP error! Status: ${res.status}`);
        }
        const html = await res.text();
        appDiv.innerHTML = html;

        //import and call js dashboard init func
        InitDashboard(logOut);

        //set css
        //remove all other css files 
        document.querySelectorAll("link[data-page]").forEach(link => link.remove());

        //create custom css as link
        const link = document.createElement("link");
        link.rel = "stylesheet";
        link.href = "/styles/dashboard.css"
        link.dataset.page = "true";
        document.head.appendChild(link);
    } else {
        //not logged in, display login
        //fetch login html & update appDiv
        const res = await fetch("/templates/login.html");
        if (!res.ok) {
            throw new Error(`HTTP error! Status: ${res.status}`);
        }
        const html = await res.text();
        console.log(html);
        appDiv.innerHTML = html;

        //set css
        //remove all other css files 
        document.querySelectorAll("link[data-page]").forEach(link => link.remove());

        //create custom css as link
        const link = document.createElement("link");
        link.rel = "stylesheet";
        link.href = "/styles/login.css"
        link.dataset.page = "true";
        document.head.appendChild(link);

        //js below
        
        //call js with the login function
        InitLogin(onLogin, onSignup);
    }
}

//auth login function
async function onLogin(name, pass) {
    try {
        const userCredential = await signInWithEmailAndPassword(auth, name, pass)
        console.log(userCredential.user)    
    } 
    catch(error) {

        let loginDiv = document.getElementById("login-error");
        loginDiv.innerHTML = 
        `
        <h4> Hatali e-posta veya sifre </h4>
        `
    }
}

//auth signup function
async function onSignup(name, pass) {
    try {
        const userCredential = await createUserWithEmailAndPassword(auth, name, pass)
        console.log(userCredential.user)    
    } 
    catch(error) {

        let loginDiv = document.getElementById("login-error");
        loginDiv.innerHTML = 
        `
        <h4> Gecersiz e-posta veya sifre </h4>
        `
    }
}

async function logOut() {
    await signOut(auth)
}