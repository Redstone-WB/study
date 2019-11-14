#### Category 1. Architectures
- Traditional ML : "bag of words" => No order
- RNN : for sequences => Maintain order
- RNN 변형 : LSTM, GRU, ... 

#### Category 2. Embeddings
- Word2Vec, GloVe
- Gives words a location
- models relationships between words
- Word similarities 
- Analogies (King - man = Queen - woman)


Machine Translation -> Seq2Seq model 보통 사용 (input seq과 output seq. 길이 달라도 ok.)

#### Word Embeddings
- Why? weights와 biases (선형결합)를 학습하여 classification/regression 을 수행하기 위해서는
input이 string이 아닌 숫자 형태여야 함. => Word Embedding
- Sol#1 : One-hot encoding
- Feature Vector ? Data table에서 each row를 가리킴
