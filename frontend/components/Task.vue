<template>
    <div class="py-2">
        <ul>
            <li v-for="task in tasks" class="py-1">
                {{task}}
                <b-button variant="primary" @click="modalShow = !modalShow,temp=task">Edit</b-button>
                <b-button variant="primary" @click="deleteTask(task)">Delete</b-button>
                <b-modal v-model="modalShow" hide-footer id="edit">
                    <b-form inline >
                        <h2>Edit Task</h2>
                        <b-form-input type="text" v-model="editTask"></b-form-input>
                        <b-button class="ml-2" @click="toggleModal" variant="primary">Edit</b-button>
                        <b-button class="ml-2" @click="$bvModal.hide('edit')" variant="primary">Close</b-button>
                     </b-form>
                </b-modal>
            </li>
        </ul>
    </div>
</template>

<script>
import axios from 'axios'
export default {
    props:{
      tasks:{
          type:Array,
          require:false
      }  
    },
    data() {
        return {
            temp: '',
            modalShow: false,
            editTask: '',
        }
    },
    methods:{
        async toggleModal() {
            this.$root.$emit('bv::toggle::modal', 'edit', 
            await axios.put('http://localhost:8000/todo',JSON.stringify({
                old_task:this.temp,
                task:this.editTask
            }), {
                headers: {'Content-Type': 'application/json'}
            })
            .then((res) => {
                console.log(res)
            })
            .catch((err) => {
                console.log(err)
            }))
        },
        async deleteTask(rs) {
            await axios.post('http://localhost:8000/tododelete',JSON.stringify({
                task:rs,
            }), {
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