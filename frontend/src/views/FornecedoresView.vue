<script setup>
import { ref, onMounted, watch } from 'vue';
import api from '@/services/api';

const fornecedores = ref([]);
const lotes = ref([]); // Lotes gerais ou filtrados
const fornecedorSelecionado = ref(null);

const formFornecedor = ref({ nome: '', cnpj: '' });

onMounted(async () => {
  const res = await api.get('/fornecedores');
  fornecedores.value = res.data;
  // Carregar todos os lotes inicialmente
  const resLotes = await api.get('/lotes');
  lotes.value = resLotes.data;
});

const criarFornecedor = async () => {
  await api.post('/fornecedores', formFornecedor.value);
  const res = await api.get('/fornecedores');
  fornecedores.value = res.data;
};

// Filtra lotes quando clica em um fornecedor
watch(fornecedorSelecionado, async (novoVal) => {
  if (!novoVal) return;
  // O endpoint /lotes aceita filtro? De acordo com o backend: filter-id_fornecedor
  const res = await api.get(`/lotes?filter-id_fornecedor=eq.${novoVal.id}`);
  lotes.value = res.data;
});
</script>

<template>
  <div class="nav-space"></div>
  <div class="container-fornecedores">
    <div class="detalhes-lotes">
      <h2>Lotes / Estoque</h2>
      <div v-for="lote in lotes" :key="lote.id_lote" class="card-lote">
        <p><strong>Item ID:</strong> {{ lote.id_produto }}</p>
        <p>Qtd: {{ lote.quantidade_inicial }} | R$ Unit: {{ lote.preco_unitario }}</p>
        <p>Validade: {{ lote.validade ? lote.validade : 'N/A' }}</p>
        <p class="warning" v-if="lote.estragados > 0">Estragados: {{ lote.estragados }}</p>
      </div>
    </div>

    <div class="lista-fornecedores">
      <h2>Fornecedores</h2>
      <div class="form-fornecedor">
        <input v-model="formFornecedor.nome" placeholder="Nome">
        <input v-model="formFornecedor.cnpj" placeholder="CNPJ">
        <button @click="criarFornecedor">+</button>
      </div>
      
      <div class="scroll-list">
        <div 
          v-for="f in fornecedores" 
          :key="f.id" 
          class="card-simple"
          @click="fornecedorSelecionado = f"
          :class="{ active: fornecedorSelecionado?.id === f.id }"
        >
          <p>{{ f.nome }}</p>
          <small>{{ f.cnpj }}</small>
        </div>
      </div>
    </div>
  </div>
</template>
