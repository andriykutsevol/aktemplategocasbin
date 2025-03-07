import React, { useState, useEffect }  from 'react'
import { CCard, CCardBody, CCol, CCardHeader, CRow } from '@coreui/react'
import {
  CChartBar,
  CChartDoughnut,
  CChartLine,
  CChartPie,
  CChartPolarArea,
  CChartRadar,
} from '@coreui/react-chartjs'


const Charts = () => {

    const [dataq, setData] = useState(null);
    const [labelsq, setLabels] = useState(null);

    const random = () => Math.round(Math.random() * 100)

    const options = {
    scales: {
        x: {
        display: false, // Hide x-axis labels
        },
        y: {
        display: false, // Hide y-axis labels
        },
    },
    plugins: {
        legend: {
        display: false, // Hide legend
        },
    },
    };

    //const url = 'http://localhost:8080/api/v1/weather/Dnipro';
    //const url = 'http://backend:8080/api/v1/weather/Dnipro';
    const url = '/api/v1/weather/Dnipro';


    const token = 'Bearer ' + window.localStorage.getItem("access_token").replace(/"/g, ''); // Replace this with your actual token

    useEffect(() => {
        const fetchData = () => {
            fetch(url, {
                method: 'GET',
                headers: {
                    'accept': 'application/json',
                    'Authorization': token,
                },    
            })
            .then(response => {
                if (!response.ok) {
                throw new Error('Network response was not ok');
                }
                return response.json(); // Returns a Promise that resolves to JSON object
            })
            .then(datao => {
                setData(datao.message.MessageJson.hourly.temperature_2m);
                setLabels(datao.message.MessageJson.hourly.time);
            })
            .catch(error => console.error('Error fetching data:', error));
        };

        fetchData();

    }, []); // Empty dependency array ensures this effect runs only once


    var dataa = {
        labels: labelsq,
        datasets: [
        {
            label: 'My First dataset',
            backgroundColor: 'rgba(220, 220, 220, 0.2)',
            borderColor: 'rgba(220, 220, 220, 1)',
            pointBackgroundColor: 'rgba(220, 220, 220, 1)',
            pointBorderColor: '#fff',
            data: dataq,
        },
        ],
    }
    return (
    <CRow>
        <CCol xs={12}>
        <CCard className="mb-4">
            <CCardHeader>Weather in Dnipro (96 days, hourly)</CCardHeader>
            <CCardBody>

            <CChartLine
                data={dataa}
                options={options}
            />

            </CCardBody>
        </CCard>
        </CCol>
    </CRow>
    ) 
}

export default Charts
