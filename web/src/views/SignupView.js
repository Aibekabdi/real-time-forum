import AbstractView from "./AbstractView.js";

function showError(err) {
    let errorDiv = document.getElementById("error-div");
    errorDiv.innerHTML = err;
}

export default class extends AbstractView{
    constructor(params) {
        super(params);
        this.setTitle("Sign up");
    }
    async getHtml(){
        return `
        <form id="signup-form" onsubmit = "return false;">
            <h3> Sign up </h3>
            
            First name: <br>
            <input type="text" id="firstname" placeholder="type your name" required>
            
            Last name: <br>
            <input type="text" id="lastname" placeholder="type your lastname" required>
            
            Email: <br>
            <input type="text" id="email" placeholder="type your email" required>
            
            Username: <br>
            <input type="text" id="nickname" placeholder="type your username" required>
            
            Age: <br>
            <input type="number" id="age" min="10" max="99" placeholder="age" required>

            Gender: <br>
            <input type="radio" name="gender" id="gender-male" value="male" required> Male
            <input type="radio" name="gender" id="gender-female" value="female"> Female <br> <br>
            
            Password: <br>
            <input type="password" id="password" placeholder="type your password"  minlength="7" maxlength="64" required>
            
            Confirm password: <br>
            <input type="password" id="psw-confirm" placeholder="Password" maxlength="64" required>
            
            <div class="error" id="error-div"></div>
            
            <button type="submit">Sign up</button>
        </form>
        `;
    }

    async init(){
        const signupForm = document.getElementById("signup-form")
        signupForm.addEventListener("submit", function(){
            let password = document.getElementById("password");
            let pswConfirm = document.getElementById("psw-confirm");
            if (password.value != pswConfirm.value){
                showError("Passwords do not match")
            }else {
                let input = {
                    email: document.getElementById("email").value,
                    nickname: document.getElementById("nickname").value,
                    firstname: document.getElementById("firstname").value,
                    lastname: document.getElementById("lastname").value,
                    password: password.value,
                    age: document.getElementById("age").value,
                    gender: document.querySelector('input[name="gender"]:checked').value,
                }
                console.log(input)
            }
        })
    }
}