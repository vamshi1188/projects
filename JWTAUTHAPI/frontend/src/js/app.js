// Constants
const API_URL = 'http://localhost:8000/api';
const app = document.getElementById('app');

// Navigation state
let isAuthenticated = false;
let currentUser = null;

// Page rendering functions
function renderNavbar() {
    const navbar = document.createElement('nav');
    navbar.className = 'navbar navbar-expand-lg navbar-dark';
    
    navbar.innerHTML = `
        <div class="container">
            <a class="navbar-brand" href="#">JWT Auth App</a>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav ms-auto">
                    ${!isAuthenticated ? `
                        <li class="nav-item">
                            <a class="nav-link" href="#" id="nav-login">Login</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="#" id="nav-register">Register</a>
                        </li>
                    ` : `
                        <li class="nav-item">
                            <a class="nav-link" href="#" id="nav-profile">Profile</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="#" id="nav-logout">Logout</a>
                        </li>
                    `}
                </ul>
            </div>
        </div>
    `;
    
    return navbar;
}

function renderLoginPage() {
    const loginForm = document.createElement('div');
    loginForm.className = 'auth-form';
    
    loginForm.innerHTML = `
        <h2>Login</h2>
        <div id="login-alert"></div>
        <form id="login-form">
            <div class="form-group">
                <label for="email">Email</label>
                <input type="email" class="form-control" id="login-email" required>
            </div>
            <div class="form-group">
                <label for="password">Password</label>
                <input type="password" class="form-control" id="login-password" required>
            </div>
            <button type="submit" class="btn btn-primary">Login</button>
        </form>
        <p class="mt-3 text-center">
            Don't have an account? <a href="#" id="go-to-register">Register</a>
        </p>
    `;
    
    return loginForm;
}

function renderRegisterPage() {
    const registerForm = document.createElement('div');
    registerForm.className = 'auth-form';
    
    registerForm.innerHTML = `
        <h2>Register</h2>
        <div id="register-alert"></div>
        <form id="register-form">
            <div class="form-group">
                <label for="name">Name</label>
                <input type="text" class="form-control" id="register-name" required>
            </div>
            <div class="form-group">
                <label for="email">Email</label>
                <input type="email" class="form-control" id="register-email" required>
            </div>
            <div class="form-group">
                <label for="password">Password</label>
                <input type="password" class="form-control" id="register-password" required>
            </div>
            <button type="submit" class="btn btn-primary">Register</button>
        </form>
        <p class="mt-3 text-center">
            Already have an account? <a href="#" id="go-to-login">Login</a>
        </p>
    `;
    
    return registerForm;
}

function renderProfilePage() {
    const profileCard = document.createElement('div');
    profileCard.className = 'profile-card';
    
    if (!currentUser) {
        profileCard.innerHTML = `
            <div class="alert alert-danger">
                You need to be logged in to view this page.
            </div>
        `;
        setTimeout(() => {
            renderPage('login');
        }, 2000);
        return profileCard;
    }
    
    profileCard.innerHTML = `
        <h2>User Profile</h2>
        <div class="profile-info">
            <strong>ID:</strong> ${currentUser.id}
        </div>
        <div class="profile-info">
            <strong>Name:</strong> ${currentUser.name}
        </div>
        <div class="profile-info">
            <strong>Email:</strong> ${currentUser.email}
        </div>
    `;
    
    return profileCard;
}

// Page rendering
function renderPage(page) {
    // Clear the app container
    app.innerHTML = '';
    
    // Add the navbar
    app.appendChild(renderNavbar());
    
    // Create container for page content
    const container = document.createElement('div');
    container.className = 'container';
    app.appendChild(container);
    
    // Render the requested page
    switch (page) {
        case 'login':
            container.appendChild(renderLoginPage());
            attachLoginListeners();
            break;
        case 'register':
            container.appendChild(renderRegisterPage());
            attachRegisterListeners();
            break;
        case 'profile':
            container.appendChild(renderProfilePage());
            break;
        default:
            if (isAuthenticated) {
                container.appendChild(renderProfilePage());
            } else {
                container.appendChild(renderLoginPage());
                attachLoginListeners();
            }
    }
    
    // Attach navbar listeners
    attachNavbarListeners();
}

// Event listeners
function attachNavbarListeners() {
    const loginLink = document.getElementById('nav-login');
    const registerLink = document.getElementById('nav-register');
    const profileLink = document.getElementById('nav-profile');
    const logoutLink = document.getElementById('nav-logout');
    
    if (loginLink) {
        loginLink.addEventListener('click', (e) => {
            e.preventDefault();
            renderPage('login');
        });
    }
    
    if (registerLink) {
        registerLink.addEventListener('click', (e) => {
            e.preventDefault();
            renderPage('register');
        });
    }
    
    if (profileLink) {
        profileLink.addEventListener('click', (e) => {
            e.preventDefault();
            renderPage('profile');
        });
    }
    
    if (logoutLink) {
        logoutLink.addEventListener('click', (e) => {
            e.preventDefault();
            logout();
        });
    }
}

function attachLoginListeners() {
    const loginForm = document.getElementById('login-form');
    const goToRegister = document.getElementById('go-to-register');
    
    loginForm.addEventListener('submit', (e) => {
        e.preventDefault();
        login();
    });
    
    goToRegister.addEventListener('click', (e) => {
        e.preventDefault();
        renderPage('register');
    });
}

function attachRegisterListeners() {
    const registerForm = document.getElementById('register-form');
    const goToLogin = document.getElementById('go-to-login');
    
    registerForm.addEventListener('submit', (e) => {
        e.preventDefault();
        register();
    });
    
    goToLogin.addEventListener('click', (e) => {
        e.preventDefault();
        renderPage('login');
    });
}

// API functions
async function login() {
    const email = document.getElementById('login-email').value;
    const password = document.getElementById('login-password').value;
    const alertBox = document.getElementById('login-alert');
    
    try {
        const response = await fetch(`${API_URL}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email, password }),
            credentials: 'include'
        });
        
        const data = await response.json();
        
        if (response.ok) {
            alertBox.innerHTML = `
                <div class="alert alert-success">
                    Login successful! Redirecting to profile...
                </div>
            `;
            
            // Get user data
            await fetchUser();
            
            setTimeout(() => {
                renderPage('profile');
            }, 1000);
        } else {
            alertBox.innerHTML = `
                <div class="alert alert-danger">
                    ${data.message || 'Login failed. Please check your credentials.'}
                </div>
            `;
        }
    } catch (error) {
        alertBox.innerHTML = `
            <div class="alert alert-danger">
                An error occurred. Please try again.
            </div>
        `;
        console.error('Login error:', error);
    }
}

async function register() {
    const name = document.getElementById('register-name').value;
    const email = document.getElementById('register-email').value;
    const password = document.getElementById('register-password').value;
    const alertBox = document.getElementById('register-alert');
    
    try {
        const response = await fetch(`${API_URL}/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ name, email, password }),
            credentials: 'include'
        });
        
        const data = await response.json();
        
        if (response.ok) {
            alertBox.innerHTML = `
                <div class="alert alert-success">
                    Registration successful! Redirecting to login...
                </div>
            `;
            
            setTimeout(() => {
                renderPage('login');
            }, 1000);
        } else {
            alertBox.innerHTML = `
                <div class="alert alert-danger">
                    ${data.message || 'Registration failed. Please try again.'}
                </div>
            `;
        }
    } catch (error) {
        alertBox.innerHTML = `
            <div class="alert alert-danger">
                An error occurred. Please try again.
            </div>
        `;
        console.error('Registration error:', error);
    }
}

async function fetchUser() {
    try {
        const response = await fetch(`${API_URL}/user`, {
            method: 'GET',
            credentials: 'include'
        });
        
        if (response.ok) {
            const userData = await response.json();
            currentUser = userData;
            isAuthenticated = true;
            return userData;
        } else {
            isAuthenticated = false;
            currentUser = null;
            return null;
        }
    } catch (error) {
        console.error('Fetch user error:', error);
        isAuthenticated = false;
        currentUser = null;
        return null;
    }
}

async function logout() {
    try {
        const response = await fetch(`${API_URL}/logout`, {
            method: 'POST',
            credentials: 'include'
        });
        
        if (response.ok) {
            isAuthenticated = false;
            currentUser = null;
            
            const container = document.querySelector('.container');
            container.innerHTML = `
                <div class="auth-form">
                    <div class="alert alert-success">
                        Logout successful! Redirecting to login...
                    </div>
                </div>
            `;
            
            setTimeout(() => {
                renderPage('login');
            }, 1000);
        }
    } catch (error) {
        console.error('Logout error:', error);
    }
}

// Initialize the app
async function initApp() {
    // Check if user is already logged in
    await fetchUser();
    
    // Render the initial page
    renderPage(isAuthenticated ? 'profile' : 'login');
}

// Start the app when the DOM is loaded
document.addEventListener('DOMContentLoaded', initApp);