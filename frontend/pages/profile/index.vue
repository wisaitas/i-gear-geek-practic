<template>
    <div class="my-3" align="center">
        <b-card
            align="center"
            title="Profile"
            tag="article"
            style="max-width: 20rem;"
            class="mb-2"
        >
            <b-img rounded="circle" :src="this.$auth.user.image_src" width=100></b-img>
            <b-card class="mt-3" align ="left">
                <b-text>
                    Name : {{this.details.first_name + " " + this.details.last_name}}
                </b-text>
            </b-card>
            <b-card align="left" class="mt-3">
                <b-text>
                    Age : {{this.details.age}}
                </b-text>
            </b-card>

            <div class="mt-4">
                <b-button @click="editProfile()" variant="primary">Edit Profiles</b-button>
            </div>
            </b-card>
    </div>
</template>

<script>
export default {
    data() {
        return {
            details:{
                first_name:this.$auth.user.first_name,
                last_name:this.$auth.user.last_name,
                age:this.$auth.user.age
            }
        }
    },
    layout: "navbar",
    middleware: 'auth',
    methods:{
        editProfile() {
            this.$router.push("/profile/edit")
        }
    }
}
</script>