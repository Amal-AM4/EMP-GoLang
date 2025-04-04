<template>
    <div>
        <h2>Update Employee</h2>
        <InsertEmployee 
            v-if="employee" 
            :id="employee.id"
            :name="employee.name" 
            :age="employee.age" 
            :position="employee.position" 
            :salary="employee.salary" 
            :isUpdate="true"
            @employeeUpdated="goBack"
        />

    </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { GetEmployee } from "../../wailsjs/go/backend/EmployeeService";
import InsertEmployee from "../components/InsertEmployee.vue";

const route = useRoute();
const router = useRouter();
const employee = ref(null);

const fetchEmployee = async () => {
    try {
        const id = parseInt(route.params.id);
        employee.value = await GetEmployee(id);
    } catch (error) {
        console.error("Error fetching employee details:", error);
    }
};

onMounted(fetchEmployee);

const goBack = () => {
    router.push("/"); // Redirect back to employee list
};


</script>
