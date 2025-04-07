const API_URL = 'http://localhost:8080/api';

async function fetchData(url, method = 'GET', data = null) {
    const options = {
        method,
        headers: {
            'Content-Type': 'application/json',
        },
    };
    
    if (data) {
        options.body = JSON.stringify(data);
    }

    try {
        const response = await fetch(url, options);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        console.error('Error:', error);
        return null;
    }
}

function createRecipeCard(recipe) {
    const recipeCard = document.createElement('div');
    recipeCard.className = 'recipe-card';
    recipeCard.innerHTML = `
        <img src="${recipe.image}" alt="${recipe.title}" class="recipe-img">
        <div class="recipe-info">
            <h3>${recipe.title}</h3>
            <p>${recipe.description}</p>
            <p><strong>Время приготовления:</strong> ${recipe.prepTime}</p>
            <a href="/recipe-detail?id=${recipe.id}">Подробнее</a>
        </div>
    `;
    return recipeCard;
}

function renderRecipeDetails(recipe) {
    const recipeDetail = document.getElementById('recipe-detail');
    recipeDetail.innerHTML = `
        <h2>${recipe.title}</h2>
        <img src="${recipe.image}" alt="${recipe.title}" class="recipe-detail-img">
        <div class="recipe-meta">
            <p><strong>Время подготовки:</strong> ${recipe.prepTime}</p>
            <p><strong>Время готовки:</strong> ${recipe.cookTime}</p>
        </div>
        <div class="recipe-content">
            <h3>Ингредиенты</h3>
            <ul id="ingredients-list">
                ${recipe.ingredients.map(ing => `<li>${ing}</li>`).join('')}
            </ul>
            <h3>Инструкции</h3>
            <ol id="instructions-list">
                ${recipe.instructions.map(inst => `<li>${inst}</li>`).join('')}
            </ol>
        </div>
        <a href="/recipes" class="btn">Вернуться к рецептам</a>
    `;
}

async function loadFeaturedRecipes() {
    const featuredRecipesContainer = document.getElementById('featured-recipes');
    if (!featuredRecipesContainer) return;

    const recipes = await fetchData(`${API_URL}/recipes`);
    if (recipes && recipes.length > 0) {
        recipes.slice(0, 3).forEach(recipe => {
            featuredRecipesContainer.appendChild(createRecipeCard(recipe));
        });
    }
}

async function loadAllRecipes() {
    const recipesContainer = document.getElementById('recipes-container');
    if (!recipesContainer) return;

    const recipes = await fetchData(`${API_URL}/recipes`);
    if (recipes && recipes.length > 0) {
        recipes.forEach(recipe => {
            recipesContainer.appendChild(createRecipeCard(recipe));
        });
    }
}

async function loadRecipeDetails() {
    const urlParams = new URLSearchParams(window.location.search);
    const recipeId = urlParams.get('id');
    
    if (!recipeId) {
        renderRecipeNotFound();
        return;
    }

    const recipe = await fetchData(`${API_URL}/recipes/${recipeId}`);
    if (recipe) {
        renderRecipeDetails(recipe);
    } else {
        renderRecipeNotFound();
    }
}

function renderRecipeNotFound() {
    const recipeDetail = document.getElementById('recipe-detail');
    if (recipeDetail) {
        recipeDetail.innerHTML = `
            <h2>Рецепт не найден</h2>
            <p>Извините, запрошенный рецепт не существует.</p>
            <a href="/recipes" class="btn">Вернуться к рецептам</a>
        `;
    }
}

function setupAddRecipeForm() {
    const form = document.getElementById('add-recipe-form');
    if (!form) return;

    document.getElementById('add-ingredient').addEventListener('click', () => {
        const div = document.createElement('div');
        div.className = 'ingredient-row';
        div.innerHTML = `
            <input type="text" name="ingredient" placeholder="Ингредиент" required>
            <input type="text" name="amount" placeholder="Количество" required>
            <button type="button" class="remove-btn">×</button>
        `;
        document.getElementById('ingredients-container').appendChild(div);
        div.querySelector('.remove-btn').addEventListener('click', () => div.remove());
    });

    document.getElementById('add-instruction').addEventListener('click', () => {
        const div = document.createElement('div');
        div.className = 'instruction-row';
        div.innerHTML = `
            <input type="text" name="instruction" placeholder="Шаг инструкции" required>
            <button type="button" class="remove-btn">×</button>
        `;
        document.getElementById('instructions-container').appendChild(div);
        div.querySelector('.remove-btn').addEventListener('click', () => div.remove());
    });

    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        
        const formData = new FormData(form);
        const ingredients = [];
        const instructions = [];
        
        document.querySelectorAll('#ingredients-container .ingredient-row').forEach(row => {
            const ingredient = row.querySelector('input[name="ingredient"]').value;
            const amount = row.querySelector('input[name="amount"]').value;
            ingredients.push(`${ingredient} - ${amount}`);
        });
        
        document.querySelectorAll('#instructions-container .instruction-row').forEach(row => {
            instructions.push(row.querySelector('input[name="instruction"]').value);
        });
        
        const recipe = {
            title: formData.get('title'),
            description: formData.get('description'),
            prepTime: `${formData.get('prep-time')} мин`,
            cookTime: `${formData.get('cook-time')} мин`,
            image: formData.get('image') || 'https://via.placeholder.com/400x300',
            ingredients,
            instructions
        };
        
        const result = await fetchData(`${API_URL}/recipes`, 'POST', recipe);
        if (result) {
            alert('Рецепт успешно добавлен!');
            form.reset();
            window.location.href = '/recipes';
        }
    });
}

document.addEventListener('DOMContentLoaded', () => {
    const path = window.location.pathname;
    
    if (path === '/' || path === '/index.html') {
        loadFeaturedRecipes();
    } else if (path === '/recipes') {
        loadAllRecipes();
    } else if (path === '/recipe-detail') {
        loadRecipeDetails();
    } else if (path === '/add-recipe') {
        setupAddRecipeForm();
    }

    document.addEventListener('click', (e) => {
        if (e.target.classList.contains('remove-btn')) {
            e.target.parentElement.remove();
        }
    });
});