<template>
    <div class="mt-2 ml-2">
        <h2>Task Board</h2>
        <p>Create a list of tasks</p>
        <b-form inline >
            <b-form-input type="text" v-model="newTask" @keypress.enter="addTask"></b-form-input>
            <b-button class="ml-2" @click="addTask" variant="primary">Add</b-button>
        </b-form>
        <Task :tasks="allTask"/>
    </div>
</template>

<script>
import axios from 'axios'
import Task from '@/components/Task'
export default {
    components:{
        Task,
    },
    data() {
        return {
            callFunc: this.getTodo(),
            newTask: '',
            allTask: [],
        }
    },
    layout: "navbar",
    methods:{
        async getTodo() {
            await axios.get('http://localhost:8000/todo')
            .then((res) => {
                for(var i in res.data){
                    this.allTask.push(JSON.parse(JSON.stringify(res.data[i])).task);
                }
            })
            .catch((err) => {
                console.log(err)
            })
        },

        async addTask(evt) {
            await axios.post('http://localhost:8000/todo',JSON.stringify({
                task:this.newTask
            }),{
                headers: {'Content-Type': 'application/json'}
            })
            .then((res) => {
                console.log(res)
            })
            .catch((err) => {
                console.log(err)
            })
        },
    }
}
</script>