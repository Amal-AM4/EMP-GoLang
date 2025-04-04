<template>
    <div class="container">
        <div v-if="employee" class="employee-card">
            <h2>Employee Details</h2>
            <div class="info">
                <p><strong>ID:</strong> {{ employee.id }}</p>
                <p><strong>Name:</strong> {{ employee.name }}</p>
                <p><strong>Age:</strong> {{ employee.age }}</p>
                <p><strong>Position:</strong> {{ employee.position }}</p>
                <p><strong>Salary:</strong> ₹{{ employee.salary.toLocaleString() }}</p>
            </div>
        </div>
        <div v-else class="loading">
            Loading employee details...
        </div>
    </div>
</template>


<script setup>
import { onMounted, ref } from "vue";
import { GetEmployee } from '../../wailsjs/go/backend/EmployeeService';
import { useRoute } from "vue-router";

const route = useRoute();

const employee = ref(null);

const fetch = async () => {
    try {
        const id = parseInt(route.params.id);
        employee.value = await GetEmployee(id);
    } catch (error) {
        alert("Error fetching employee: " + error);
    }
};

onMounted(fetch);

</script>

<style lang="scss" scoped>
.container {
    max-width: 600px;
    margin: 40px auto;
    padding: 20px;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;

    .employee-card {
        background: #ffffff;
        border-radius: 12px;
        box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
        padding: 24px 32px;
        color: #333;

        h2 {
            margin-bottom: 20px;
            color: #2E6049;
            font-size: 24px;
            text-align: center;
        }

        .info {
            p {
                font-size: 18px;
                padding: 8px 0;
                border-bottom: 1px solid #f0f0f0;
                text-transform: capitalize;

                strong {
                    color: #555;
                    margin-right: 8px;
                }
            }
        }
    }

    .loading {
        text-align: center;
        color: #777;
        font-size: 18px;
        margin-top: 60px;
    }
}
</style>