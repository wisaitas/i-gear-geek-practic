<template>
    <div>
        <b-navbar type="dark" variant="dark">
            <b-navbar-nav>
                <nuxt-link class="nav-item nav-link" to="/">Home</nuxt-link>
                <div v-if="$auth.loggedIn">
                    <div @click="$auth.logout()"><nuxt-link class="nav-item nav-link" to="/">Logout</nuxt-link></div>
                </div>
                <div v-else>
                    <nuxt-link class="nav-item nav-link" to="/login">Login</nuxt-link>
                </div>
                <nuxt-link class="nav-item nav-link" to="/todo">Todolist</nuxt-link>
                <nuxt-link class="nav-item nav-link" to="/register">Register</nuxt-link>

                <b-nav-item-dropdown text="Profile" left>
                    <b-dropdown-item href="/profile">Profile</b-dropdown-item>
                    <b-dropdown-item href="/profile/edit">Edit Profile</b-dropdown-item>
                </b-nav-item-dropdown>
            </b-navbar-nav>
        </b-navbar>
        <nuxt/>
    </div>
</template>
