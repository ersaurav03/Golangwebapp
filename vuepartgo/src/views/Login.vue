<template>
  <div>
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a
          role="button"
          class="navbar-burger burger"
          aria-label="menu"
          aria-expanded="false"
          data-target="navbarBasicExample"
        >
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>
      <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item">
            <router-link to="/">Home</router-link>
          </a>
          <a class="navbar-item">
            <router-link to="/about">About</router-link>
          </a>
          <a class="navbar-item">
            <router-link to="/contact">Contact</router-link>
          </a>
          <a class="navbar-item">
            <router-link to="/login">Login</router-link>
          </a>
          <a class="navbar-item">
            <router-link to="/register">Register</router-link>
          </a>
        </div>
      </div>
    </nav>
    <section class="hero is-success is-fullheight">
      <div class="hero-body">
        <div class="container">
          <div class="columns">
            <div class="column">
              <h1 class="title">Login page</h1>
              <!-- <h2 class="subtitle">Login Page</h2> -->
            </div>
            <div class="column"></div>
            <div class="column"></div>
            <div class="column">
              <div class="field">
                <label class="label">Email Id</label>
                <div class="control">
                  <input class="input" type="email" placeholder="Email">
                </div>
              </div>
              <div class="field">
                <label class="label">Password</label>
                <div class="control">
                  <input class="input" type="password" placeholder="Pasword">
                </div>
              </div>

              <div class="field is-grouped">
                <div class="control">
                  <button class="button is-link">SignIn</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          <strong>Change URL</strong> by
          <a href="#">JSMPJ Corporation</a>. The source code is licensed
          <a href="#">MIT</a>. The website content
          is licensed
          <a href="#">CC BY NC SA 4.0</a>.
        </p>
      </div>
    </footer>
  </div>
</template>
