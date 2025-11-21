<script setup>
import { ref } from 'vue';
import api from '@/services/api';

const periodo = ref({ start: '', end: '' });
const dadosRelatorio = ref(null);

const gerarRelatorio = async () => {
  if(!periodo.value.start || !periodo.value.end) return alert("Selecione as datas");
  
  try {
    const res = await api.get('/relatorios/financeiro', {
      params: {
        start: periodo.value.start,
        end: periodo.value.end,
        granularity: 'month'
      }
    });
    dadosRelatorio.value = res.data;
  } catch (e) {
    console.error(e);
  }
};
</script>

<template>
  <div class="nav-space"></div>
  <div class="relatorio-container">
    <h1>Contabilidade</h1>
    
    <div class="filtros">
      <label>Início: <input type="date" v-model="periodo.start"></label>
      <label>Fim: <input type="date" v-model="periodo.end"></label>
      <button @click="gerarRelatorio">Gerar</button>
    </div>

    <div v-if="dadosRelatorio" class="resultado">
      <div class="totais">
        <div class="card red">
          <h3>Despesa</h3>
          <p>R$ {{ dadosRelatorio.totals.despesa.toFixed(2) }}</p>
        </div>
        <div class="card green">
          <h3>Receita</h3>
          <p>R$ {{ dadosRelatorio.totals.receita.toFixed(2) }}</p>
        </div>
        <div class="card yellow">
          <h3>Lucro/Saldo</h3>
          <p>R$ {{ dadosRelatorio.totals.lucro.toFixed(2) }}</p>
        </div>
      </div>

      <table>
        <thead>
          <tr><th>Data</th><th>Receita</th><th>Despesa</th><th>Lucro</th></tr>
        </thead>
        <tbody>
          <tr v-for="ponto in dadosRelatorio.series" :key="ponto.date">
            <td>{{ ponto.date }}</td>
            <td>{{ ponto.receita }}</td>
            <td>{{ ponto.despesa }}</td>
            <td>{{ ponto.lucro }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.card { padding: 20px; border-radius: 8px; color: white; font-weight: bold;}
.red { background-color: #e74c3c; }
.green { background-color: #27ae60; }
.yellow { background-color: #f1c40f; color: black; }
.totais { display: flex; gap: 20px; margin-top: 20px; }
</style>
