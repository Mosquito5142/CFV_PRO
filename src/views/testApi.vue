<template>
    <button @click="testApi" class="mb-4 p-2 bg-red-500 text-white rounded-md">Test API</button>
</template>

<script setup>
import axios from 'axios';

const testApi = () => {
    const testDeckData = {
        name: 'Test Deck',
        cards: ['card1', 'card2', 'card3']
    };

    axios.post('/api/macros/s/AKfycbwxQSc_FdDlkU81payets_7BpBtqs61yDI_yLcIe_of7sArG3fhg-UZps8VIW3s486FvA/exec', testDeckData, {
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => {
            console.log('Success:', response.data);
            alert('Test data uploaded successfully!');
        })
        .catch((error) => {
            console.error('Error:', error);
            let errorMessage = 'Failed to upload test data.';

            if (error.response) {
                // The request was made and the server responded with a status code
                // that falls out of the range of 2xx
                console.error('Error data:', error.response.data);
                console.error('Error status:', error.response.status);
                console.error('Error headers:', error.response.headers);
                errorMessage += ` Server responded with status: ${error.response.status}. Data: ${JSON.stringify(error.response.data)}`;
            } else if (error.request) {
                // The request was made but no response was received
                console.error('Error request:', error.request);
                errorMessage += ' No response received from server.';
                console.error('Request details:', error.request);
            } else {
                // Something happened in setting up the request that triggered an Error
                console.error('Error message:', error.message);
                errorMessage += ` Message: ${error.message}`;
            }

            alert(errorMessage);
        });
};
</script>