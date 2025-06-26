# math-maps

### 1. Sistema de Coordenadas Geograficas
Referencial: Linha do Equador
( φ ) Latitude -> -90º (Sul) | +90º (Norte)
( λ ) Longitude -> -180º (Oeste) | +180º  (Leste)

### 2. Cercos Poligonais
Centro = [φ₀ | λ₀] -> Respectivamente latitude e longitude do ponto central

L = Lado do quadrado (km)
Δφ = (L/2) / 111 -> Deslocamento em latitude
Δλ = (L/2) / (111 * cos(φ₀)) -> // em longitude

Canto Superior Esquerdo: (φ₀ + Δφ, λ₀ - Δλ)
Canto Superior Direito: (φ₀ + Δφ, λ₀ + Δλ)
Canto Inferior Esquerdo: (φ₀ - Δφ, λ₀ - Δλ)
Canto Inferior Direito: (φ₀ - Δφ, λ₀ + Δλ)

e.g
Definir variaveis principais:
Centro do Quadrado: (φ₀ = -23.55 | λ₀ = -46.63)
Lado do Quadrado: L = 2km

Descobrir o deslocamento em graus da lat e lon:
Δφ = (2 / 2) / 111 ≈ 0.0090° -> Precisa pegar os graus por conta da curvatura da terra, se não da merda
Δλ = (2 / 2) / (111 × cos(23.55°)) ≈ 0.0098° -> Basicamente a mesma coisa mas adiciona o coseno da latitude

Vértices Especificos:
CS Esquerdo (-23.55 + 0.0090, -46.63 - 0.0098) ≈ (-23.541, -46.640)
CS Direito (-23.55 + 0.0090, -46.63 + 0.0098) ≈ (-23.541, -46.620)
CI Esquerdo (-23.55 - 0.0090, -46.63 - 0.0098) ≈ (-23.559, -46.640)
CI Direito (-23.55 - 0.0090, -46.63 + 0.0098) ≈ (-23.559, -46.620)

Para saber se um determinado ponto está dentro ou fora desse quadrado, precisamos estabelecer algumas coisas antes, primeiro, o quadrado DEVE ter angulos de 90º, já que vamos utilizar um algoritmo de bounding box simples, segundo, precisamos pegar algumas constantes, como a menor / maior valor de latitude / longitude. Aí conseguiremos aplicar o binding box:

Ponto a ser verificado: P(φ, λ) -> P(-23.541, -46.620)

IF
φ >= φ_min
φ <= φ_max
λ >= λ_min
λ <= λ_max
{print("ta dentro da box")}
Se esses valores se aplicarem ao ponto, ele pode estar dentro do quadrado