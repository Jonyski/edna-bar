<script setup>
import { ref, onMounted, computed } from 'vue';
import api from '@/services/api';

// Estado
const clientes = ref([]);
const funcionarios = ref([]);
const produtos = ref([]);
const carrinho = ref([]);

const vendaForm = ref({
  id_cliente: null,
  id_funcionario: null,
  tipo_pagamento: 'pix' // Default
});

// Carregar dados iniciais
onMounted(async () => {
  try {
    const [resCli, resFunc, resProd] = await Promise.all([
      api.get('/clientes'),
      api.get('/funcionarios'),
      api.get('/produtos/comercial') // Apenas produtos vendáveis
    ]);
    clientes.value = resCli.data;
    funcionarios.value = resFunc.data;
    produtos.value = resProd.data;
  } catch (error) {
    console.error("Erro ao carregar dados", error);
  }
});

// Ações
const adicionarAoCarrinho = (produto) => {
  const itemExistente = carrinho.value.find(item => item.id === produto.id);
  if (itemExistente) {
    itemExistente.quantidade++;
  } else {
    carrinho.value.push({ ...produto, quantidade: 1 });
  }
};

const totalVenda = computed(() => {
  return carrinho.value.reduce((acc, item) => acc + (item.preco_venda * item.quantidade), 0);
});

const finalizarVenda = async () => {
  try {
    // 1. Cria a venda
    const resVenda = await api.post('/vendas', {
      id_cliente: vendaForm.value.id_cliente,
      id_funcionario: vendaForm.value.id_funcionario,
      tipo_pagamento: vendaForm.value.tipo_pagamento,
      data_hora_renda: new Date().toISOString(), // Backend espera data_hora_renda (Go struct tag)
      data_hora_pagamento: new Date().toISOString()
    });

    const idVenda = resVenda.data.id;

    // 2. Cria os itens da venda (Backend endpoint: /item_venda)
    // Nota: Seu backend espera IDLote. Isso é complexo para o front decidir (FIFO).
    // Idealmente o backend deveria resolver o lote automaticamente. 
    // Por enquanto, assumiremos que o backend trataria isso ou teríamos que buscar lotes.
    alert(`Venda ${idVenda} realizada com sucesso! Total: R$ ${totalVenda.value}`);
    carrinho.value = [];
  } catch (error) {
    alert("Erro ao finalizar venda: " + error.response?.data?.detail || error.message);
  }
};
</script>

<template>
  <div class="nav-space"></div>
  <div class="container-vendas">
    <div class="painel-venda">
      <h2>Nova Venda</h2>
      <div class="form-group">
        <label>Cliente</label>
        <select v-model="vendaForm.id_cliente">
          <option v-for="c in clientes" :key="c.id" :value="c.id">{{ c.nome }}</option>
        </select>
      </div>
      <div class="form-group">
        <label>Funcionário</label>
        <select v-model="vendaForm.id_funcionario">
          <option v-for="f in funcionarios" :key="f.id" :value="f.id">{{ f.nome }}</option>
        </select>
      </div>
      
      <div class="carrinho-lista">
        <div v-for="item in carrinho" :key="item.id">
          {{ item.nome }} x{{ item.quantidade }} - R$ {{ (item.preco_venda * item.quantidade).toFixed(2) }}
        </div>
      </div>
      
      <div class="total">
        <h3>Total: R$ {{ totalVenda.toFixed(2) }}</h3>
        <button @click="finalizarVenda">Finalizar ($)</button>
      </div>
    </div>

    <div class="grid-produtos">
      <div v-for="prod in produtos" :key="prod.id" class="card-produto" @click="adicionarAoCarrinho(prod)">
        <h4>{{ prod.nome }}</h4>
        <p>{{ prod.marca }}</p>
        <p class="preco">R$ {{ prod.preco_venda }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Adicione seu CSS aqui para alinhar colunas conforme o desenho */
.container-vendas { display: flex; gap: 20px; }
.grid-produtos { display: grid; grid-template-columns: repeat(auto-fill, minmax(150px, 1fr)); gap: 10px; width: 60%; }
.card-produto { border: 1px solid #ccc; padding: 10px; cursor: pointer; }
</style>
