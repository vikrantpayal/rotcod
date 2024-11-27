# rotcod
Exploring applications of AIML and IOT in the medical field. Uses htmx, golang.

python --version #3.8 or higher
python -m venv env
pip install boto3 moviepy
sudo apt update;sudo apt install ffmpeg;

# CONFIGURE aws  cli
sudo apt update
sudo apt upgrade -y
sudo apt install -y curl unzip
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
aws --version
aws configure
# enter key id and access key
export PATH=$PATH:/usr/local/bin
echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc
source ~/.bashrc
# TEST
aws s3 ls


